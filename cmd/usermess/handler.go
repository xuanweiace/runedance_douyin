package main

import (
	"context"
	"runedance_douyin/kitex_gen/usermess"
)

// UsermessServiceImpl implements the last service interface defined in the IDL.
type UsermessServiceImpl struct{}

// UserRegister implements the UsermessServiceImpl interface.
func (s *UsermessServiceImpl) UserRegister(ctx context.Context, req *usermess.DouyinUserRegisterRequest) (resp *usermess.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...

	return
}

// UserLogin implements the UsermessServiceImpl interface.
func (s *UsermessServiceImpl) UserLogin(ctx context.Context, req *usermess.DouyinUserLoginRequest) (resp *usermess.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...

	return
}

// GetUserMess implements the UsermessServiceImpl interface.
func (s *UsermessServiceImpl) GetUserMess(ctx context.Context, req *usermess.DouyinUserRequest) (resp *usermess.DouyinUserResponse, err error) {
	// TODO: Your code here...

	return
}
