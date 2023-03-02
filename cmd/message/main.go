package main

import (
	"log"
	
	message "runedance_douyin/kitex_gen/message/messageservice"
	"runedance_douyin/cmd/message/dal"
	
	"net"
	constants "runedance_douyin/pkg/consts"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"

	etcd "github.com/kitex-contrib/registry-etcd"
)


func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	dal.Init()

	svr := message.NewServer(new(MessageServiceImpl),
	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MessageServiceName}),
	server.WithServiceAddr(&net.TCPAddr{Port: constants.MessageServicePort}),
	server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
	server.WithMuxTransport(),
	server.WithRegistry(r),
	)
	
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}






	// svr := message.NewServer(new(MessageServiceImpl))
	// dal.Init()
	// err := svr.Run()
	// if err != nil {
	// 	log.Println(err.Error())
	// }
}

