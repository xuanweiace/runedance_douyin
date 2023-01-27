package dal

import "runedance_douyin/cmd/user/dal/db_mysql"

func Init() {
	db_mysql.MySQLInit()

}
