package main

import (
	"log"
	"runedance_douyin/cmd/usermess/dal"
	usermess "runedance_douyin/kitex_gen/usermess/usermessservice"
)

func main() {
	dal.Init()
	svr := usermess.NewServer(new(UsermessServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
