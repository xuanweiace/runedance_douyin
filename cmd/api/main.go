// Code generated by hertz generator.

package main

import (
	"runedance_douyin/cmd/api/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.Init()
}
func main() {
	Init()
	// 127.0.0.1
	h := server.Default(server.WithHostPorts("0.0.0.0:8080"))
	//h.Use(mw.MyJWT())

	register(h)
	h.Spin()
}
