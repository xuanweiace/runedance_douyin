package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"runedance_douyin/cmd/interaction/dal"
	"runedance_douyin/cmd/interaction/rpc"
	interaction "runedance_douyin/kitex_gen/interaction/interactionservice"
	constants "runedance_douyin/pkg/consts"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	dal.Init()
	rpc.Init()
	svr := interaction.NewServer(new(InteractionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.InteractionServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: constants.InteractionServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}