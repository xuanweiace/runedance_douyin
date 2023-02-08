package constants

const (
	//db_mysql
	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

	//user
	JWTSecketKey    = "runedance"
	UserServiceName = "user_service"
	UserServicePort = 9001

	//relation
	RelationTableName         = "relation"
	RelationServiceName       = "relation_service"
	RelationServicePort       = 9000
	ActionType_AddRelation    = 1
	ActionType_RemoveRelation = 2

	//message
	MessageTableName = "message"
	MessageServiceName = "message_service"
	MessageServicePort       = 9002

	//etcd
	EtcdAddress = "127.0.0.1:2379"
)
