package main

import (
	"log"
	"net"
	"runedance_douyin/cmd/relation/dal"
	"runedance_douyin/cmd/relation/rpc"
	relation "runedance_douyin/kitex_gen/relation/relationservice"

	"github.com/cloudwego/kitex/server"
)

func Init() {
	rpc.Init()
	dal.Init()
}
func main() {

	Init()
	svr := relation.NewServer(new(RelationServiceImpl), server.WithServiceAddr(&net.TCPAddr{Port: 9000}))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
