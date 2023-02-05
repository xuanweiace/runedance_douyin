package runedance_douyin

import (
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"runedance_douyin/pkg/middleware"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/login", middleware.JwtMiddleware.LoginHandler)

	// your code ...
}
