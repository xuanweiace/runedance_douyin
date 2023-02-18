package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	recommend "runedance_douyin/cmd/recommend/kitex_gen/recommend/recommendservice"
)

var gormClient *gorm.DB

func getdsn() string {
	return "root:xtr1593jT@tcp(qwq.bogo.ac.cn:3306)/back?charset=utf8mb4&parseTime=True&loc=Local"
}
func main() {

	db, e1 := gorm.Open(mysql.Open(getdsn()), &gorm.Config{})
	if e1 != nil {
		panic(e1)
		return
	}
	gormClient = db

	svr := recommend.NewServer(new(RecommendServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
