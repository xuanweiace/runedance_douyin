package consts

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"runedance_douyin/pkg/errnos"
)

var (
	//db_mysql
	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

	//relation
	RelationTableName   = "relation"
	RelationServiceName = "relation_service"
	//videos
	VideoTableName string
	//user
	UserTableName string
	//etcd
	EtcdAddress = "127.0.0.1:2379"
)

func Init() {
	work, err := os.Getwd()
	if err != nil {
		str := "获取错误路径出现问题"
		errnos.Wrap(err, str)
	}
	fmt.Println(work)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work)
	MySQLDefaultDSN = viper.GetString("DSN")
	RelationTableName = viper.GetString("Table.relation")
	VideoTableName = viper.GetString("Table.videos")
	UserTableName = viper.GetString("Table.user")
}
