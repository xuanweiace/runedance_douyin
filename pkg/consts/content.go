package consts

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"runedance_douyin/pkg/errnos"
)

var (
	MySQLDefaultDSN string = "root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	//relation
	RelationTableName   string = "relation"
	RelationServiceName string = "relation_service"

	//etcd
	EtcdAddress string = "127.0.0.1:2379"
)

func readInfo() {
	work, err := os.Getwd()
	if err != nil {
		str := "获取错误路径出现问题"
		errnos.Wrap(err, str)
	}
	fmt.Println(work)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work)
	err = viper.ReadInConfig()
	if err != nil {
		errnos.Wrap(err, "读取配置文件问题")
	}
	account := viper.GetString("DSN.account")
	password := viper.GetString("DSN.password")
	dbname := viper.GetString("DSN.Dbname")
	host := viper.GetString("DSN.host")
	//读取正式内容
	MySQLDefaultDSN = account + ":" + password + "@tcp(localhost:" + host + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	RelationTableName = viper.GetString("Table.relation")
	RelationServiceName = viper.GetString("ServiceName")
	EtcdAddress = viper.GetString("etcd.address")

}
