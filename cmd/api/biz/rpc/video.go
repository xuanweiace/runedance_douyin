package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	videoprocess "runedance_douyin/kitex_gen/videoProcess"
	"runedance_douyin/kitex_gen/videoProcess/videoprocessservice"
	constants "runedance_douyin/pkg/consts"
	"time"
)

var videoProcessClient videoprocessservice.Client

func initVideo() {
	r1, err4 := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err4 != nil {
		panic(err4)
		return
	}
	c, err5 := videoprocessservice.NewClient(constants.VideoProcessServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(time.Second),
		client.WithResolver(r1),
	)
	if err5 != nil {
		panic(err5)
		return
	}
	videoProcessClient = c
}
func PublishVideo(author int64, title string, fileData []byte) (int32, string) {
	resp, err := videoProcessClient.UploadVideo(context.TODO(), &videoprocess.VideoProcessUploadRequest{
		AuthorId:  author,
		Title:     title,
		VideoData: fileData,
	})
	if resp.Result_ == false || err != nil {
		fmt.Println(resp.Message)
		fmt.Println(err)
		return 1002, "上传失败" //
	} else {
		return 1001, "ok"
	}
}
