package dal

import (
	"runedance_douyin/cmd/message/dal/db_redis"
)

func Init(){
	db_redis.Init()
}