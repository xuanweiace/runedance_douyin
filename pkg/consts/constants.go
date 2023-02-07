package constants

const (
	//db_mysql
	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

	//relation
	RelationTableName   = "relation"
	RelationServiceName = "relation_service"

	//message
	MessageTableName = "message"
	MessageServiceName = "message_service"

	//etcd
	EtcdAddress = "127.0.0.1:2379"
)
