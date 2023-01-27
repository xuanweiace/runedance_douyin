package main

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
		str := ""
		errnos.Wrap(err, str)
	}
	viper.SetConfigName("info")
	viper.SetConfigFile("yml")
	fmt.Println(work)
}
