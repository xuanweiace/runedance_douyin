package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"runedance_douyin/cmd/relation/dal"
	relation "runedance_douyin/kitex_gen/relation/relationservice"
)

func main() {
	dal.Init()

	svr := relation.NewServer(new(RelationServiceImpl), server.WithServiceAddr(&net.TCPAddr{Port: 9000}))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
