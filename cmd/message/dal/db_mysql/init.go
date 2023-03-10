package db_mysql

import (
	// "gorm.io/sharding"

	"log"
	"runedance_douyin/pkg"
	// constants "runedance_douyin/pkg/consts"
	
	"gorm.io/gorm"
)

var db *gorm.DB

// Init init DB
func Init() {
	// constants.MySQLDefaultDSN
	pkg.InitDB("root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local")
	db = pkg.GetDB()

	isExist := db.Migrator().HasTable(&MessageRecord{})
	if(!isExist){
		err := db.Migrator().CreateTable(&MessageRecord{})
		if err != nil {
			log.Printf("fail to create table :%s\n", err)
			return
		}
	}

	// sharding
	// db.Use(sharding.Register(sharding.Config{
	// 	ShardingKey: "user_to_user",
	// 	NumberOfShards: 64,
	// 	PrimaryKeyGenerator: sharding.PKSnowflake,
	// }, constants.MessageTableName))

	// if err := db.Use(tracing.NewPlugin()); err != nil {
	// 	panic(err)
	// }

}