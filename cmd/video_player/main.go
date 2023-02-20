// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	constants "runedance_douyin/pkg/consts"
	"strconv"
)

func main() {
	server.WithHostPorts("0.0.0.0:" + strconv.Itoa(constants.VideoPlayUrlPort))
	h := server.Default()
	register(h)
	h.Spin()
}
