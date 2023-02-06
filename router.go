package runedance_douyin

import (
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_jwt/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
