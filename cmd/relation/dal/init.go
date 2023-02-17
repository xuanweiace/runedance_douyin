package dal

import (
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/cmd/relation/dal/db_redis"
)

func Init() {
	db_mysql.Init()

	//todo  redis
	db_redis.Init()
}
