package main

import (
	"log"
	usermess "runedance_douyin/kitex_gen/usermess/usermessservice"
)

func main() {
	svr := usermess.NewServer(new(UsermessServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
