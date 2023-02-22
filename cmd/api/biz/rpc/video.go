package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	log "github.com/sirupsen/logrus"
	"runedance_douyin/cmd/api/biz/model/douyin"
	"runedance_douyin/kitex_gen/interaction"
	"runedance_douyin/kitex_gen/user"
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
func PublishVideo(ctx context.Context, author int64, title *[]byte, fileData *[]byte) (int32, string) {
	resp, err := videoProcessClient.UploadVideo(ctx, &videoprocess.VideoProcessUploadRequest{
		AuthorId:  author,
		Title:     string(*title),
		VideoData: *fileData,
	})
	if err != nil || resp.Result_ == false {
		return 1002, "上传失败"
	} else {
		return 0, "ok"
	}
}
func GetUser(ctx context.Context, requesterID int64, userID int64) (*douyin.User, string, error) {
	//uc, _ := userservice.NewClient("127.0.0.1:1234")
	u, err1 := userClient.GetUser(context.TODO(), &user.DouyinUserRequest{UserId: userID, MyUserId: requesterID})
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
func GetVideo(ctx context.Context, requesterID int64, videoID int64) (*douyin.Video, string, error) {
	v, err2 := videoProcessClient.GetVideoInfo(ctx, &videoprocess.VideoInfoRequest{VideoId: videoID})
	if err2 != nil {
		return nil, "获取视频信息失败", err2
	}
	u, msg, err1 := GetUser(ctx, requesterID, v.AuthorId)
	if err1 != nil {
		return nil, "获取用户详细信息失败:" + msg, err1
	}
	f, err := interactionClient.GetFavoriteStatus(ctx, &interaction.GetFavoriteStatusRequest{VideoId: videoID, UserId: requesterID})
	if err != nil {
		return nil, "获取点赞关系失败", err
	}
	var IsFavorite bool
	if f.ActionType == 1 {
		IsFavorite = true
	} else {
		IsFavorite = false
	}
	return &douyin.Video{
		ID:            v.VideoId,
		Author:        u,
		PlayURL:       constants.VideoUrlPrefix + v.StorageId,
		CoverURL:      constants.CoverUrlPrefix + v.StorageId,
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount:  int64(v.CommentCount),
		IsFavorite:    IsFavorite,
		Title:         v.Title,
	}, "ok", nil
}
func GetPublishList(ctx context.Context, requesterID int64, authorID int64) ([]*douyin.Video, string, error) {
	videoIDList, err := videoProcessClient.GetVideoList(ctx, authorID)
	if err != nil {
		return nil, "获取发布列表失败", err
	}
	if len(videoIDList.Published) < 1 {
		return nil, "发布列表为空", nil
	}
	var resp []*douyin.Video
	for _, vid := range videoIDList.Published {
		//if err != nil {
		//	continue
		//	//return nil, "", err
		//}
		//info.
		v, msg, e := GetVideo(ctx, requesterID, vid)
		if e != nil {
			log.WithFields(log.Fields{
				"user_id":    requesterID,
				"request_id": ctx.Value("request_id"),
			}).Warning("信息获取失败，跳过该条视频:" + msg)
			continue
		} else {
			resp = append(resp, v)
		}
	}
	return resp, "ok", nil
}
func GetRecommendList(ctx context.Context, requesterID int64) ([]*douyin.Video, string, error) {
	videoIDList, err := videoProcessClient.GetVideoList(ctx, 0)
	if err != nil || videoIDList == nil || len(videoIDList.Published) < 1 {
		return nil, "推荐列表获取失败", err
	}
	var resp []*douyin.Video
	for _, vid := range videoIDList.Published {
		v, msg, e := GetVideo(ctx, requesterID, vid)
		if e != nil {
			log.WithFields(log.Fields{
				"user_id":    requesterID,
				"request_id": ctx.Value("request_id"),
			}).Warning("信息获取失败，跳过该条视频:" + msg)
			continue
		} else {
			resp = append(resp, v)
		}
	}
	return resp, "ok", nil
}
