package constants

const (
	//db_mysql
	MySQLDefaultDSN = "root:mysqlmm200107@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

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

	VideoStorageServiceName = "video_storage_service"
	VideoProcessServiceName = "video_process_service"
	VideoStorageServicePort = 9003
	VideoProcessServicePort = 9002
	VideoUrlSuffix          = "_transcode_100030.mp4"
	CoverUrlSuffix          = "_snapshotByOffset_10_0.jpg"
	CoverUrlPrefix          = "http://127.0.0.1:8888/fileapi/cover/"
	VideoUrlPrefix          = "http://127.0.0.1:8888/fileapi/video/"

	//etcd
	EtcdAddress = "127.0.0.1:2379"

	//interaction
	InteractionTableName   = "interaction"
	InteractionServiceName = "interaction_service"
	InteractionServicePort = 9006
)
