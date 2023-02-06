package main

import (
	"log"
	"runedance_douyin/cmd/user/dal"
	user "runedance_douyin/kitex_gen/user/userservice"
)

func main() {
	dal.Init()
	svr := user.NewServer(new(UserServiceImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
