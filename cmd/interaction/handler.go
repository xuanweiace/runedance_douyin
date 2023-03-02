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
type InteractionServiceImpl struct{}

const (
	timeLayout = "2006-01-02 15:04:05"
)

// FavoriteAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteRequest) (resp *interaction.FavoriteResponse, err error) {
	resp = interaction.NewFavoriteResponse()
	//更新mysql
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
	} else {
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
	//1 放缓存 favorite
	key := "fav" + strconv.FormatInt(req.UserId, 10) + strconv.FormatInt(req.VideoId, 10)
	updateFavoriteToRedis(ctx, req.ActionType, key)
	var fc int32
	if req.ActionType == 1 {
		fc = updateVideo.FavoriteCount + 1
	} else {
		fc = updateVideo.FavoriteCount - 1
	}
	err = rpc.ChangeFavoriteCountToDB(req.VideoId, fc)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		return resp, err
	}
	return resp, err
}

// GetFavoriteList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetFavoriteList(ctx context.Context, req *interaction.GetFavoriteListRequest) (resp *interaction.GetFavoriteListResponse, err error) {
	resp = interaction.NewGetFavoriteListResponse()
	//1 从redis中查询fl
	vidLists, err := mysql.GetFavoriteDao().GetFavoriteList(req.UserId)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	for _, vid := range vidLists {
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
		userMeta, err := rpc.GetUser(video.AuthorId)
		if err != nil {
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		var user *interaction.User
		user = &interaction.User{
			UserId:        userMeta.UserId,
			Username:      userMeta.Username,
			FollowCount:   userMeta.FollowCount,
			FollowerCount: userMeta.FollowerCount,
			IsFollow:      userMeta.IsFollow,
		}
		resp.VedioList = append(resp.VedioList, &interaction.Video{
			Id:            vid,
			Author:        user,
			PlayUrl:       constants.VideoUrlPrefix + video.StorageId,
			CoverUrl:      constants.VideoUrlPrefix + video.StorageId,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    true,
			Title:         video.Title,
		})

	}
	return resp, err
}

// CommentAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentRequest) (resp *interaction.CommentResponse, err error) {
	resp = interaction.NewCommentResponse()

	now := time.Now()
	if req.ActionType == 1 {
		log.Println(tools.FilterSensitive(req.GetCommentText()))
		comment := &mysql.Comment{
			Id:           time.Now().UnixNano() / 1000000,
			Uid:          req.UserId,
			Vid:          req.VideoId,
			Content:      tools.FilterSensitive(req.GetCommentText()),
			Content_date: now.Unix() / 1000,
		}
		err = mysql.GetCommentDao().AddComment(comment)
		if err != nil {
			log.Println("add fail")
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		userMeta, err := rpc.GetUser(comment.Uid)
		if err != nil {
			log.Println("getUser fail")
			resp.StatusCode = errnos.CodeServiceErr
			er := err.Error()
			resp.StatusMsg = &er
			return resp, err
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
		resp.Comment = &interaction.Comment{
			Id: comment.Id,
			User: &interaction.User{
				UserId:        userMeta.UserId,
				Username:      userMeta.Username,
				FollowCount:   userMeta.FollowCount,
				FollowerCount: userMeta.FollowerCount,
				IsFollow:      userMeta.IsFollow,
			},
			Content:    tools.FilterSensitive(req.GetCommentText()),
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
		log.Println("getVideo fail")
		log.Println(req.VideoId)
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
	err = rpc.ChangeCommentCountToDB(req.VideoId, cc)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
		return resp, err
	}
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
	key := "fav" + strconv.FormatInt(req.UserId, 10) + strconv.FormatInt(req.VideoId, 10)
	status, _ := queryFavoriteFromRedis(ctx, key)
	if status == 0 {
		favorite, err := mysql.GetFavoriteDao().FindFavorite(req.UserId, req.VideoId)
		if err != nil {
			resp.StatusCode = 0
			resp.StatusMsg = nil
			resp.ActionType = 2
			return resp, nil
		}
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.ActionType = favorite.Action
		updateFavoriteToRedis(ctx, resp.ActionType, key)
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.ActionType = status
	}
	return resp, nil
}
