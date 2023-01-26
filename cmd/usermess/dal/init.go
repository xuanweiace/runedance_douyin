package dal

import "runedance_douyin/cmd/usermess/dal/db_mysql"

func Init() {
	db_mysql.MySQLInit()

}
