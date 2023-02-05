package middleware

import (
	"context"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/dal/mysql"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"github.com/pkg/errors"
	"net/http"
	"runedance_douyin/pkg/tools"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)
var loginStruct struct {
	username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
} //接收用户信息结构体
func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("douyin"), //todo
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		//用于设置授权已认证的用户路由访问权限的函数
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := mysql.CheckUser(loginStruct.username, tools.Md5Util(loginStruct.Password, "test")) //todo
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user not exists or wrong password")
			}
			return users[0], nil //user[0]提供payload数据源
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		//成功的话返回值
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": code,
				"status_msg":  "success",
				"user_id":     loginStruct.username,
				"token":       token,
			})
		}, //返回值
		IdentityKey: IdentityKey,
		//用于设置获取身份信息的函数,配合idkey存取让上下文使用
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				UserName: claims[IdentityKey].(string),
			}
		},
		//jwt token校验
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		//验证失败
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":       -1,
				"status_msg": "验证失败",
				"user_id":    loginStruct.username,
				"token":      "-1",
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
