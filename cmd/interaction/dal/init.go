package dal

import (
	"runedance_douyin/pkg"
)

func Init() {
	pkg.InitDB("root:mysqlmm200107@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local")

}
