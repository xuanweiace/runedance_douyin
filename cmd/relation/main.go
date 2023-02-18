package main

import (
	"log"
	"net"
	"runedance_douyin/cmd/relation/dal"
	"runedance_douyin/cmd/relation/rpc"
	relation "runedance_douyin/kitex_gen/relation/relationservice"
	constants "runedance_douyin/pkg/consts"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	rpc.Init()
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	Init()
	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: constants.RelationServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
