package dal

import (
	"runedance_douyin/cmd/message/dal/db_redis"
)

func init(){
	db_redis.Init()
}