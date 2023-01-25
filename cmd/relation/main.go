package main

import (
	"log"
	"runedance_douyin/cmd/relation/dal"
	relation "runedance_douyin/kitex_gen/relation/relationservice"
)

func main() {
	dal.Init()

	svr := relation.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
