package dal

import (
	"runedance_douyin/pkg"
	constants "runedance_douyin/pkg/consts"
)

func Init() {

	pkg.InitDB(constants.MySQLDefaultDSN)

}
