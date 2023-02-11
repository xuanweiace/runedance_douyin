package main

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"net/url"
	videoprocess "runedance_douyin/kitex_gen/videoProcess/videoprocessservice"
	"runedance_douyin/kitex_gen/videoStorage/videostorageservice"
	constants "runedance_douyin/pkg/consts"
	"time"
)

var cosClient *cos.Client
var gormClient *gorm.DB
var storageClient videostorageservice.Client

func getdsn() string {
	return constants.MySQLDefaultDSN
	//return "root:1@tcp(qwq.bogo.ac.cn:3306)/back?charset=utf8mb4&parseTime=True&loc=Local"
}
func main() {

	db, err1 := gorm.Open(mysql.Open(getdsn()), &gorm.Config{})
	if err1 != nil {
		panic(err1)
		return
	}
	gormClient = db

	u, _ := url.Parse("https://bogo-1308981928.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	cc := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDBBxWq4D8oDTaT7UeieQ1TYTDj15wb5LP",
			SecretKey: "ymOp6BmCPOAi291gJn6FD0azDF5jBRVO",
		},
	})
	if cc == nil {
		panic("gg")
		return
	}
	cosClient = cc

	r1, err4 := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err4 != nil {
		panic(err4)
		return
	}
	vsc, err5 := videostorageservice.NewClient(constants.VideoStorageServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(time.Second),
		client.WithResolver(r1),
	)
	if err5 != nil {
		panic(err5)
		return
	}
	storageClient = vsc

	//vsc, e2 := videostorageservice.NewClient("VideoStorageService", client.WithHostPorts("0.0.0.0:8888"))
	//if e2 != nil {
	//	panic(e2)
	//	return
	//}
	//storageClient = vsc

	r2, err3 := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err3 != nil {
		panic(err3)
		return
	}
	svr := videoprocess.NewServer(new(VideoProcessServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoProcessServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: constants.VideoProcessServicePort}),
		server.WithMuxTransport(),
		server.WithRegistry(r2),
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         200,
			UpdateControl:  nil,
		}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
