package main

import (
	"log"
	"net"
	"runedance_douyin/cmd/user/dal"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	user "runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	dal.Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: constants.UserServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
