package main

import (
	"log"
	message "runedance_douyin/kitex_gen/message/messageservice"
	"runedance_douyin/cmd/message/dal"
	// "runedance_douyin/cmd/message/middleware"

)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))
	dal.Init()
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
	// middleware.Init()
}
