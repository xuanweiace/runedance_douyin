package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	// 可根据需要自行添加字段
	Username string `json:"username"`
	User_id  int64  `json:"user_id"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("runedance")

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
type Response struct {
	StatusCode int32
	StatusMsg  string
}

func MyJWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("进入token鉴权")
		auth := c.Query("token")
		if len(auth) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "无token字段")
			return
		}
		token, err := ParseToken(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "鉴权失败")
			return
		} else {
			println("token 正确")
			//不知道后续怎么用加了个tokenflag
			c.Set("token_f", 1)
			c.Set("username", token.Username)
			c.Set("user_id", token.User_id)
			c.Next(ctx)
		}

		//不知道后续怎么用加了个tokenflag
		c.Set("token_f", 1)
		c.Set("username", token.Username)
		c.Set("user_id", token.User_id)

		c.Next(ctx)

	}
}
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/*func MyJWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
			Realm:      "test",
			Key:        []byte("douyin"),
			Timeout:    time.Hour,
			MaxRefresh: time.Hour,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*User); ok {
					return jwt.MapClaims{
						identityKey: v.UserName,
					}
				}
				return jwt.MapClaims{}
			},
			IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
				claims := jwt.ExtractClaims(ctx, c)
				return &User{
					UserName: claims[identityKey].(string),
				}
			},
			Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
				var loginVals login
				if err := c.BindAndValidate(&loginVals); err != nil {
					return "", jwt.ErrMissingLoginValues
				}
				userID := loginVals.Username
				password := loginVals.Password
				if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
					return &User{
						UserName:  userID,
						LastName:  "Hertz",
						FirstName: "CloudWeGo",
					}, nil
				}
				return nil, jwt.ErrFailedAuthentication
			},
			Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
				if v, ok := data.(*User); ok && v.UserName == "admin" {
					return true
				}

				return false
			},
			Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
				c.JSON(code, map[string]interface{}{
					"code":    code,
					"message": message,
				})
			},
		})
		if err != nil {
			log.Fatal("JWT Error:" + err.Error())
		}

		// When you use jwt.New(), the function is already automatically called for checking,
		// which means you don't need to call it again.
		errInit := authMiddleware.MiddlewareInit()

		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
		c.Next(ctx)
	}
}*/
