package main

import (
	"log"
	"runedance_douyin/cmd/interaction/dal"
	interaction "runedance_douyin/kitex_gen/interaction/messageservice"
)

func main() {
	dal.Init()
	svr := interaction.NewServer(new(InteractionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
