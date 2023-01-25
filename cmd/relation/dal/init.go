package dal

import "runedance_douyin/cmd/relation/dal/db_mysql"

func Init() {
	db_mysql.Init()

	//todo  redis
}
