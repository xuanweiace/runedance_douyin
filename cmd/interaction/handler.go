package main

import (
	"context"
	"log"
	"runedance_douyin/cmd/interaction/dal/mysql"
	"runedance_douyin/cmd/interaction/rpc"
	"runedance_douyin/kitex_gen/interaction"
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
	// TODO: Your code here...
	//video表的FavoriteCount+1
	//1 验证token登陆有效
	resp = interaction.NewFavoriteResponse()

	//2 查出记录 没有则插入记录
	favorite, err := mysql.GetFavoriteDao().FindFavorite(req.UserId, req.VideoId)
	if err != nil {
		//需插入数据
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
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	return resp, err
}

// GetFavoriteList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetFavoriteList(ctx context.Context, req *interaction.GetFavoriteListRequest) (resp *interaction.GetFavoriteListResponse, err error) {
	// TODO: Your code here...
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
		//1 由vid查出video
		//2 由AuthorId 查出user
		userMeta, _ := rpc.GetUser(1)
		var user *interaction.User
		user = &interaction.User{
			UserId:        userMeta.UserId,
			Username:      userMeta.Username,
			FollowCount:   userMeta.FollowCount,
			FollowerCount: userMeta.FollowerCount,
			IsFollow:      userMeta.IsFollow,
		}
		//由vid得到vedio list进行append !!!!!!!!!!RPC

		vedioList = append(vedioList, &interaction.Video{
			Id:            vid,
			Author:        user,
			PlayUrl:       "url",
			CoverUrl:      "url",
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true,
			Title:         "视频" + strconv.Itoa(int(vid)),
		})
		resp.VedioList = vedioList
	}

	return resp, err
}

// CommentAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentRequest) (resp *interaction.CommentResponse, err error) {
	// TODO: Your code here...
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
		} else {
			resp.StatusCode = 0
			resp.StatusMsg = nil
		}
	}
	return resp, err
}

// GetCommentList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetCommentList(ctx context.Context, req *interaction.GetCommentListRequest) (resp *interaction.GetCommentListResponse, err error) {
	// TODO: Your code here...

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
		//由vid得到vedio list进行append !!!!!!!!!!RPC
		resp.CommentList = append(resp.CommentList, &interaction.Comment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: time.Unix(comment.Content_date*1000, 0).Format(timeLayout),
		})
	}
	return resp, err
}
