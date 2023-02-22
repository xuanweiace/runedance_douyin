package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	videoprocess "runedance_douyin/kitex_gen/videoProcess"
	"runedance_douyin/kitex_gen/videoProcess/videoprocessservice"
	constants "runedance_douyin/pkg/consts"
	"time"
)

var videoClient videoprocessservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := videoprocessservice.NewClient(constants.VideoProcessServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond), // 50ms会超时
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideo(vid int64) (*videoprocess.VideoInfoResponse, error) {
	req := videoprocess.VideoInfoRequest{
		VideoId: vid,
	}
	resp, err := videoClient.GetVideoInfo(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	return resp, err
}

func ChangeFavoriteCountToDB(vid int64, num int32) error {
	req := videoprocess.ChangeCountRequest{
		VideoId:     vid,
		ChangeValue: num,
	}
	err := videoClient.ChangeFavCount(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		return err
	}
	return nil
}

func ChangeCommentCountToDB(vid int64, num int32) error {
	req := videoprocess.ChangeCountRequest{
		VideoId:     vid,
		ChangeValue: num,
	}
	err := videoClient.ChangeComCount(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		return err
	}
	return nil
}
