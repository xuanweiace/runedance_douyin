// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"runedance_douyin/cmd/api/biz/rpc"
)

func Init() {
	rpc.Init()
}
func main() {
	Init()
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	//h.Use(mw.MyJWT())

	register(h)
	h.Spin()
}
