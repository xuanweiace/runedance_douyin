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
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	//随机生成salt
	salt := tools.RandomStringUtil()
	username := req.Username
	//密码MD5加密
	password := tools.Md5Util(req.Password, salt)
	token := username + password
	//更新用户ID
	userIdSequence = db_mysql.GetUserService().FindLastUserId()
	//注册用户
	var msg string
	if err := db_mysql.GetUserService().Register(username, password, userIdSequence, salt); err != nil {
		//注册失败返回错误信息
		resp.UserId = 0
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//成功注册
		atomic.AddInt64(&userIdSequence, 1)
		resp.StatusCode = success
		resp.UserId = userIdSequence
		resp.Token = token
		msg = "Register successfully"
		resp.StatusMsg = &msg
	}
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	username := req.Username
	password := req.Password
	// todo token部分暂时未完成
	token := username + password
	//登录验证失败
	//返回：msg:user does not exist | password error
	var msg string
	if userResp, err := db_mysql.GetUserService().UserLogin(username, password); err != nil {
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//登陆成功
		resp.StatusCode = success
		msg = "Successful login"
		resp.StatusMsg = &msg
		resp.UserId = userResp.UserId
		resp.Token = token
	}
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}
