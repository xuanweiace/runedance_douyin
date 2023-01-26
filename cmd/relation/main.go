package main

import (
	"log"
	"runedance_douyin/cmd/relation/dal"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	relation "runedance_douyin/kitex_gen/relation/relationservice"
)

func main() {
	dal.Init()
	err := db_mysql.DeleteRelation(&db_mysql.Relation{
		FansID: 1,
		UserID: 2,
	})

	svr := relation.NewServer(new(RelationServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
