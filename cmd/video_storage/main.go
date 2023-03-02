package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"net/http"
	net_url "net/url"
	videostorage "runedance_douyin/kitex_gen/videoStorage/videostorageservice"
	constants "runedance_douyin/pkg/consts"
)

func getVideoDB() (string, error) {
	u := "mongodb://user02:User02@qwq.bogo.ac.cn:23317/fdb02"
	return u, nil
}

var mongoClient *mongo.Client
var cosClient *cos.Client

func main() {

	url, err1 := getVideoDB()
	c, err2 := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err1 != nil || err2 != nil {
		panic("gg")
	}
	mongoClient = c
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			//TODO 打印日志并报告数据库关闭失败
		}
	}()
	url2, _ := net_url.Parse("https://bogo-1308981928.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: url2}
	cc := cos.NewClient(b, &http.Client{})
	if cc == nil {
		panic("gg")
		return
	}
	cosClient = cc

	r, err3 := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err3 != nil {
		panic(err3)
		return
	}
	svr := videostorage.NewServer(new(VideoStorageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoStorageServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: constants.VideoStorageServicePort}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
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
