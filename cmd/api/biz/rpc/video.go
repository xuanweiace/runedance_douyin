package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"runedance_douyin/cmd/api/biz/model/douyin"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
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
func PublishVideo(author int64, title *[]byte, fileData *[]byte) (int32, string) {
	resp, err := videoProcessClient.UploadVideo(context.TODO(), &videoprocess.VideoProcessUploadRequest{
		AuthorId:  author,
		Title:     string(*title),
		VideoData: *fileData,
	})
	if resp.Result_ == false || err != nil {
		fmt.Println(resp.Message) //TODO log
		fmt.Println(err)
		return 1002, "上传失败" //TODO 错误码怎么设？
	} else {
		return 0, "ok"
	}
}
func GetUser(requesterID int64, userID int64) (*douyin.User, string, error) {
	uc, _ := userservice.NewClient("127.0.0.1:1234") //TODO etcd获取client
	u, err1 := uc.GetUser(context.TODO(), &user.DouyinUserRequest{UserId: userID, MyUserId: requesterID})
	if err1 != nil {
		return nil, *u.StatusMsg, err1
	}
	return &douyin.User{
		ID:            u.User.UserId,
		Name:          u.User.Username,
		FollowCount:   u.User.FollowCount,
		FollowerCount: u.User.FollowerCount,
		IsFollow:      u.User.IsFollow,
	}, *u.StatusMsg, nil
}
func GetVideo(requesterID int64, videoID int64) (*douyin.Video, string, error) {
	v, err2 := videoProcessClient.GetVideoInfo(context.TODO(), &videoprocess.VideoInfoRequest{VideoId: videoID})
	if err2 != nil {
		return nil, "获取视频信息失败", err2
	}
	u, msg, err1 := GetUser(requesterID, v.AuthorId)
	if err1 != nil {
		return nil, "获取用户详细信息失败:" + msg, err1
	}
	//TODO 是否点赞
	return &douyin.Video{
		ID:            v.VideoId,
		Author:        u,
		PlayURL:       constants.VideoUrlPrefix + v.StorageId,
		CoverURL:      constants.CoverUrlPrefix + v.StorageId,
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount:  int64(v.CommentCount),
		IsFavorite:    false,
		Title:         v.Title,
	}, "ok", nil
}
func GetPublishList(requesterID int64, authorID int64) ([]*douyin.Video, string, error) {
	videoIDList, err := videoProcessClient.GetVideoList(context.TODO(), authorID)
	if err != nil {
		return nil, "获取失败", err
	}
	if len(videoIDList.Published) < 1 {
		return nil, "列表为空", nil
	}
	var resp []*douyin.Video
	for _, vid := range videoIDList.Published {
		//if err != nil {
		//	continue
		//	//return nil, "", err
		//}
		//info.
		v, msg, e := GetVideo(requesterID, vid)
		if e != nil {
			fmt.Println(msg) //TODO
			continue
		} else {
			resp = append(resp, v)
		}
	}
	return resp, "ok", nil
}
