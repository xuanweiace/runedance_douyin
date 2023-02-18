package dal

import (
	"runedance_douyin/cmd/user/dal/db_mysql"
	"runedance_douyin/cmd/user/rpc"
)

func Init() {
	rpc.Init()
	db_mysql.MySQLInit()

}
