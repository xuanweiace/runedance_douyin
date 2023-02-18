// Code generated by hertz generator.
package douyin

import (
	"context"
	"io"
	douyin "runedance_douyin/cmd/api/biz/model/douyin"
	pack "runedance_douyin/cmd/api/biz/pack"
	"runedance_douyin/cmd/api/biz/rpc"
	"runedance_douyin/pkg/errnos"
	
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationAction .
// @router /douyin/rpc/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {

	var err error
	var req douyin.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(douyin.RelationActionResponse)
	if err := rpc.RelationAction(ctx, c.GetInt64("user_id"), req.ToUserID, req.ActionType); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFollowList .
// @router /douyin/rpc/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetFollowListResponse)
	if userList, err := rpc.GetFollowList(ctx, req.UserID); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		resp.UserList = pack.ConvertUserlist(userList)
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/rpc/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetFollowerListResponse)

	if userList, err := rpc.GetFollowerList(ctx, req.UserID); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		resp.UserList = pack.ConvertUserlist(userList)
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/rpc/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetFriendListResponse)

	if userList, err := rpc.GetFriendList(ctx, req.UserID); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		resp.UserList = pack.ConvertUserlist(userList)
	}

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.MessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.MessageActionResponse)

	if err := rpc.MessageAction(ctx, c.GetInt64("user_id"), req.ToUserID, req.Content); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	}

	c.JSON(consts.StatusOK, resp)
}

// GetMessageChat .
// @router /douyin/message/chat/ [POST]
func GetMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetMessageChatResponse)

	msgList, err := rpc.GetMessageChat(ctx, c.GetInt64("user_id"), req.ToUserID)
	if(err != nil){
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	}

	var result []*douyin.Message
	
	for _, val:= range msgList {
		msg := douyin.Message{
			ID: val.Id,
			Content: val.Content,
			CreateTime: &val.CreateTime,
		}
		result = append(result, &msg)
	} 
	
	resp.MsgList = result
	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUser .
// @router /douyin/user/ [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req douyin.DouyinPublishActionRequest
	//err = c.BindAndValidate(&req)
	// check token there?
	var id int64
	id = c.GetInt64("user_id")
	file, err1 := c.FormFile("data")
	fileOpen, err2 := file.Open()
	tt := c.FormValue("title")
	if err1 != nil || len(tt) < 6 || err2 != nil || file.Size > 998244353 {
		c.String(consts.StatusBadRequest, "bad request")
		return
	}
	fileData, _ := io.ReadAll(fileOpen)
	resp := new(douyin.DouyinPublishActionResponse)
	var msg string
	resp.StatusCode, msg = rpc.PublishVideo(id, &tt, &fileData)
	resp.StatusMsg = &msg
	c.JSON(200, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, msg, e := rpc.GetPublishList(c.GetInt64("user_id"), req.UserID)

	resp := new(douyin.DouyinPublishListResponse)

	resp.StatusMsg = &msg
	if e != nil {
		resp.StatusCode = 1
	} else {
		resp.StatusCode = 0
		resp.VideoList = list
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.FavoriteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.FavoriteResponse)
	if err := rpc.FavoriteAction(ctx, c.GetInt64("user_id"), req.VideoID, req.ActionType); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
	}

	c.JSON(consts.StatusOK, resp)
}

// GetFavoriteList .
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetFavoriteListResponse)
	if videoList, err := rpc.GetFavoriteList(ctx, c.GetInt64("user_id")); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		if videoList != nil {
			resp.VedioList = pack.ConvertVideolist(videoList)
		}

	}

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.CommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.CommentResponse)
	if comment, err := rpc.CommentAction(ctx, c.GetInt64("user_id"), req.VideoID, req.ActionType, *req.CommentText, *req.CommentID); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		if comment != nil {
			user := &douyin.User{
				ID:            comment.User.UserId,
				Name:          comment.User.Username,
				FollowCount:   comment.User.FollowerCount,
				FollowerCount: comment.User.FollowerCount,
			}
			resp.Comment = &douyin.Comment{
				ID:         comment.Id,
				User:       user,
				Content:    comment.Content,
				CreateDate: comment.CreateDate,
			}
		}
	}
	c.JSON(consts.StatusOK, resp)
	return
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.GetCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin.GetCommentListResponse)

	if commentList, err := rpc.GetCommentList(ctx, req.VideoID); err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = errnos.CodeSuccess
		resp.StatusMsg = nil
		if commentList != nil {
			resp.CommentList = pack.ConvertCommentlist(commentList)
		}

	}
	c.JSON(consts.StatusOK, resp)
	return
}
