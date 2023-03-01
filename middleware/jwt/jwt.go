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

func MyJWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("c.URI:", c.URI())
		fmt.Println("进入token鉴权")
		auth, ok := c.GetQuery("token")
		if ok == false {
			auth, ok = c.GetPostForm("token")
		}
		if ok == false {
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
