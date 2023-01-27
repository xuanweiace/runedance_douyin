package main

import (
	"context"
	"runedance_douyin/cmd/user/dal/db_mysql"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/pkg/tools"
	"sync/atomic"
)

const (
	success = 0
	Failed  = 1
)

var userIdSequence int64

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	//todo 账户密码校验
	var msg string
	resp := user.NewDouyinUserRegisterResponse()
	//随机生成salt
	salt := tools.RandomStringUtil()
	username := req.Username
	//密码MD5加密
	password := tools.Md5Util(req.Password, salt)
	//更新用户ID
	userIdSequence = db_mysql.GetUserService().FindLastUserId()
	//注册用户
	err := db_mysql.GetUserService().UserRegister(username, password, salt)
	if err != nil {
		//注册失败返回错误信息
		resp.UserId = 0
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//成功注册
		msg = "UserRegister successfully"
		atomic.AddInt64(&userIdSequence, 1)
		token, err := tools.GenToken(username, userIdSequence)
		if err != nil {
			msg = "token generation failed" + err.Error()
			resp.StatusMsg = &msg
			resp.StatusCode = Failed
			return resp, err
		}
		resp.StatusCode = success
		resp.UserId = userIdSequence
		resp.Token = token
		resp.StatusMsg = &msg
	}
	return resp, err
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	// todo 账户密码校验
	username := req.Username
	password := req.Password
	//登录验证失败
	//返回：msg:user does not exist | password error
	var msg string
	var resp = user.NewDouyinUserLoginResponse()
	userResp, err := db_mysql.GetUserService().UserLogin(username, password)
	if err != nil {
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//登陆成功
		resp.StatusCode = success
		msg = "Successful login"
		resp.StatusMsg = &msg
		resp.UserId = userResp.UserId
		token, _ := tools.GenToken(username, resp.UserId)
		//todo 错误处理
		resp.Token = token
	}
	return resp, err
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	var msg string
	var resp = user.NewDouyinUserResponse()
	claims, err := tools.ParseToken(req.Token)
	if err == nil { //鉴权是否登录
		resp.User, err = db_mysql.GetUserService().GetUserById(req.UserId, claims.User_id)
		if err == nil {
			resp.StatusCode = success
			msg = "GetUser success"
			resp.StatusMsg = &msg
		} else {
			msg = "get user information failed"
		}

	} else {
		msg = "login failed"
	}
	resp.StatusCode = Failed
	resp.StatusMsg = &msg
	return resp, err
}
