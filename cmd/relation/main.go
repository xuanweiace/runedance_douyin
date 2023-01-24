package main

import (
	"log"
	relation "runedance_douyin/kitex_gen/relation/relationservice"
)

func main() {
	svr := relation.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
