package main

import (
	"context"
	"log"
	"runedance_douyin/cmd/interaction/dal/mysql"
	"runedance_douyin/cmd/interaction/rpc"
	"runedance_douyin/kitex_gen/interaction"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"runedance_douyin/pkg/tools"
	"strconv"
	"time"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct {
}

const (
	timeLayout = "2006-01-02 15:04:05"
)

// FavoriteAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteRequest) (resp *interaction.FavoriteResponse, err error) {
	resp = interaction.NewFavoriteResponse()
	favorite, err := mysql.GetFavoriteDao().FindFavorite(req.UserId, req.VideoId)
	if err != nil {
		favorite = &mysql.Favorite{
			Id:     tools.RandomStringUtil(),
			Uid:    req.UserId,
			Vid:    req.VideoId,
			Action: req.ActionType,
		}
		err = mysql.GetFavoriteDao().AddFavorite(favorite)
		if err != nil {
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		return resp, err
	}
	err = mysql.GetFavoriteDao().UpdateFavorite(favorite.Id, req.ActionType)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		return resp, err
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	updateVideo, err := rpc.GetVideo(req.VideoId)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		return resp, err
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	var fc int32
	if req.ActionType == 1 {
		fc = updateVideo.FavoriteCount + 1
	} else {
		fc = updateVideo.FavoriteCount - 1
	}
	//TODO update video表
	return resp, err
}

// GetFavoriteList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetFavoriteList(ctx context.Context, req *interaction.GetFavoriteListRequest) (resp *interaction.GetFavoriteListResponse, err error) {
	//1 验证token登陆有效
	resp = interaction.NewGetFavoriteListResponse()
	//2 通过uid查出vidlist
	//var vidLists []int64
	vidLists, err := mysql.GetFavoriteDao().GetFavoriteList(req.GetUserId())
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	log.Println(vidLists)
	//3 通过vid查出vediolist
	//var vedioList []Ve
	vedioList := []*interaction.Video{}
	for _, vid := range vidLists {
		userMeta, _ := rpc.GetUser(1)
		var user *interaction.User
		user = &interaction.User{
			UserId:        userMeta.UserId,
			Username:      userMeta.Username,
			FollowCount:   userMeta.FollowCount,
			FollowerCount: userMeta.FollowerCount,
			IsFollow:      userMeta.IsFollow,
		}
		video, err := rpc.GetVideo(vid)
		if err != nil {
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		favorite, err := mysql.GetFavoriteDao().FindFavorite(req.UserId, vid)
		var flag bool
		if favorite.Action != 1 {
			flag = false
		} else {
			flag = true
		}
		vedioList = append(vedioList, &interaction.Video{
			Id:            vid,
			Author:        user,
			PlayUrl:       constants.VideoUrlPrefix + video.StorageId,
			CoverUrl:      constants.VideoUrlPrefix + video.StorageId,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    flag,
			Title:         video.Title,
		})

	}
	resp.VedioList = vedioList
	return resp, err
}

// CommentAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentRequest) (resp *interaction.CommentResponse, err error) {
	//video表的CommentCount+1
	//1 验证token登陆有效
	//var msg string
	resp = interaction.NewCommentResponse()

	now := time.Now()
	if req.ActionType == 1 {
		comment := &mysql.Comment{
			Id:           time.Now().UnixNano() / 1000000,
			Uid:          req.UserId,
			Vid:          req.VideoId,
			Content:      req.GetCommentText(),
			Content_date: now.Unix() / 1000,
		}
		err = mysql.GetCommentDao().AddComment(comment)
		if err != nil {
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		userMeta, _ := rpc.GetUser(comment.Uid)
		resp.Comment = &interaction.Comment{
			Id: comment.Id,
			User: &interaction.User{
				UserId:        userMeta.UserId,
				Username:      userMeta.Username,
				FollowCount:   userMeta.FollowCount,
				FollowerCount: userMeta.FollowerCount,
				IsFollow:      userMeta.IsFollow,
			},
			Content:    req.GetCommentText(),
			CreateDate: time.Unix(now.Unix(), 0).Format(timeLayout),
		}
	} else if req.ActionType == 2 {
		//注意error
		cid, _ := strconv.ParseInt(req.GetCommentId(), 10, 64)
		err = mysql.GetCommentDao().DeleteComment(cid)
		if err != nil {
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
	}
	updateVideo, err := rpc.GetVideo(req.VideoId)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		return resp, err
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	var cc int32
	if req.ActionType == 1 {
		cc = updateVideo.CommentCount + 1
	} else {
		cc = updateVideo.CommentCount - 1
	}
	//TODO update video表
	return resp, err
}

// GetCommentList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetCommentList(ctx context.Context, req *interaction.GetCommentListRequest) (resp *interaction.GetCommentListResponse, err error) {
	resp = interaction.NewGetCommentListResponse()

	commentList, err := mysql.GetCommentDao().FindCommentListByVid(req.VideoId)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	for _, comment := range commentList {
		userMeta, _ := rpc.GetUser(comment.Uid)
		var user *interaction.User
		user = &interaction.User{
			UserId:        userMeta.UserId,
			Username:      userMeta.Username,
			FollowCount:   userMeta.FollowCount,
			FollowerCount: userMeta.FollowerCount,
			IsFollow:      userMeta.IsFollow,
		}
		resp.CommentList = append(resp.CommentList, &interaction.Comment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: time.Unix(comment.Content_date*1000, 0).Format(timeLayout),
		})
	}
	return resp, err
}

// GetFavoriteStatus implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) GetFavoriteStatus(ctx context.Context, req *interaction.GetFavoriteStatusRequest) (resp *interaction.GetFavoriteStatusResponse, err error) {
	resp = interaction.NewGetFavoriteStatusResponse()
	favorite, err := mysql.GetFavoriteDao().FindFavorite(req.UserId, req.VideoId)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		resp.ActionType = 2
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.ActionType = favorite.Action
	return resp, err
}
