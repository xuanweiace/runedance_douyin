package main

import (
	"context"
	"log"
	"runedance_douyin/cmd/interaction/dal/mysql"
	"runedance_douyin/kitex_gen/interaction"
	"runedance_douyin/pkg/tools"
	"strconv"
	"time"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct {
}

// FavoriteAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteRequest) (resp *interaction.FavoriteResponse, err error) {
	// TODO: Your code here...
	//1 验证token登陆有效
	var msg string
	resp = interaction.NewFavoriteResponse()
	claims, err := tools.ParseToken(req.Token)
	//var uid=claims.User_id
	if err == nil { //鉴权是否登录  !!!!!!!!!测试使用
		print(err)
		resp.StatusCode = 0 //success
		msg = "success"
		resp.StatusMsg = &msg
		claims.User_id = 1
	} else {
		msg = "login invalid"
		resp.StatusCode = 1 //failed
		resp.StatusMsg = &msg
		return resp, err
	}

	//2 查出记录 没有则插入记录
	favorite, err := mysql.GetFavoriteDao().FindFavorite(claims.User_id, req.VideoId)
	if err != nil {
		//需插入数据
		favorite = &mysql.Favorite{
			Id:     tools.RandomStringUtil(),
			Uid:    claims.User_id,
			Vid:    req.VideoId,
			Action: req.ActionType,
		}
		mysql.GetFavoriteDao().AddFavorite(favorite)
		return resp, nil
	}
	mysql.GetFavoriteDao().UpdateFavorite(favorite.Id, req.ActionType)
	return resp, err

}

// GetFavoriteList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetFavoriteList(ctx context.Context, req *interaction.GetFavoriteListRequest) (resp *interaction.GetFavoriteListResponse, err error) {
	// TODO: Your code here...
	//1 验证token登陆有效
	var msg string
	resp = interaction.NewGetFavoriteListResponse()
	claims, err := tools.ParseToken(req.Token)
	//var uid=claims.User_id
	log.Println("列表err")
	log.Println(err)
	if err == nil { //鉴权是否登录  !!!!!!!!!测试使用
		print(err)
		resp.StatusCode = 0 //success
		msg = "success"
		resp.StatusMsg = &msg
		//claims.User_id = 1
	} else {
		msg = "login invalid"
		resp.StatusCode = 1 //failed
		resp.StatusMsg = &msg
		return resp, err
	}
	//2 通过uid查出vidlist
	//var vidLists []int64
	vidLists, err := mysql.GetFavoriteDao().GetFavoriteList(claims.User_id)
	log.Println(vidLists)
	//3 通过vid查出vediolist
	//var vedioList []Ve
	for _, vid := range vidLists {
		var user *interaction.User
		user = &interaction.User{
			UserId:   1,
			Username: "作者",
			IsFollow: true,
		}
		//由vid得到vedio list进行append !!!!!!!!!!RPC
		resp.VedioList = append(resp.VedioList, &interaction.Video{
			Id:            vid,
			Author:        user,
			PlayUrl:       "url",
			CoverUrl:      "url",
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true,
			Title:         "视频" + strconv.Itoa(int(vid)),
		})
	}

	return resp, err
}

// CommentAction implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentRequest) (resp *interaction.CommentResponse, err error) {
	// TODO: Your code here...
	//1 验证token登陆有效
	//var msg string
	resp = interaction.NewCommentResponse()
	claims, err := tools.ParseToken(req.Token)
	//var uid=claims.User_id
	if err == nil { //鉴权是否登录  !!!!!!!!!测试使用
		print(err)
		resp.StatusCode = 0 //success
		resp.StatusMsg = &interaction.CommentResponse_StatusMsg_DEFAULT
		claims.User_id = 1
	} else {
		var msg *string
		msg = new(string)
		*msg = "login invalid"
		resp.StatusCode = 1 //failed
		resp.StatusMsg = msg
		return resp, err
	}
	//2 插入 1或者删除 2评论
	/*
		Id           int64
			Uid          int64
			Vid          int64
			Content      string
			Content_date time.Time*/
	now := time.Now()
	timeLayout := "2006-01-02 15:04:05"
	if req.ActionType == 1 {
		comment := &mysql.Comment{
			Id:           time.Now().UnixNano() / 1000000,
			Uid:          claims.User_id,
			Vid:          req.VideoId,
			Content:      req.GetCommentText(),
			Content_date: now.Unix() / 1000,
		}
		mysql.GetCommentDao().AddComment(comment)
		//!!!!!!!!!!RPC'
		resp.Comment = &interaction.Comment{
			Id: comment.Id,
			User: &interaction.User{
				UserId:   1,
				Username: "作者",
				IsFollow: true,
			},
			Content:    req.GetCommentText(),
			CreateDate: time.Unix(now.Unix(), 0).Format(timeLayout),
		}
		//resp.Comment.Id = comment.Id
		//resp.Comment.Content = req.GetCommentText()
		//resp.Comment.User = nil
	} else if req.ActionType == 2 {
		//注意error
		mysql.GetCommentDao().DeleteComment(req.GetCommentId())
	}
	return resp, err
}

// GetCommentList implements the MessageServiceImpl interface.
func (s *InteractionServiceImpl) GetCommentList(ctx context.Context, req *interaction.GetCommentListRequest) (resp *interaction.GetCommentListResponse, err error) {
	// TODO: Your code here...
	var msg string
	resp = interaction.NewGetCommentListResponse()
	claims, err := tools.ParseToken(req.Token)
	//var uid=claims.User_id
	if err == nil { //鉴权是否登录  !!!!!!!!!测试使用
		print(err)
		resp.StatusCode = 0 //success
		msg = "success"
		resp.StatusMsg = &msg
		claims.User_id = 1
	} else {
		msg = "login invalid"
		resp.StatusCode = 1 //failed
		resp.StatusMsg = &msg
		return resp, err
	}
	commentList, _ := mysql.GetCommentDao().FindCommentListByVid(req.VideoId)
	for _, comment := range commentList {
		var user *interaction.User
		user = &interaction.User{
			UserId:   1,
			Username: "作者",
			IsFollow: true,
		}
		//由vid得到vedio list进行append !!!!!!!!!!RPC
		resp.CommentList = append(resp.CommentList, &interaction.Comment{
			Id:         comment.Id,
			User:       user,
			Content:    "123",
			CreateDate: "123",
		})
	}
	return resp, err
}
