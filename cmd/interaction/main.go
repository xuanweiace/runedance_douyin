package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"runedance_douyin/cmd/interaction/dal"
	interaction "runedance_douyin/kitex_gen/interaction/messageservice"
)

func main() {
	dal.Init()
	svr := interaction.NewServer(new(InteractionServiceImpl), server.WithServiceAddr(&net.TCPAddr{Port: 9000}))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
