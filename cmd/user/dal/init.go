package dal

import (
	"runedance_douyin/cmd/user/dal/db_mysql"
	"runedance_douyin/cmd/user/dal/db_redis"
	"runedance_douyin/cmd/user/rpc"
)

func Init() {
	rpc.Init()
	db_mysql.MySQLInit()
	db_redis.InitRedis("redis://43.143.130.52:6379", "123456")
}
