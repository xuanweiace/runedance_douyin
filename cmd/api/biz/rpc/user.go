package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(constants.UserServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond), // 50ms会超时
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	userClient = c

}

// UserRegister implements the UserServiceImpl interface.
func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	return userClient.UserRegister(ctx, req, callopt.WithRPCTimeout(3*time.Second))
}

// UserLogin implements the UserServiceImpl interface.
func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	return userClient.UserLogin(ctx, req, callopt.WithRPCTimeout(3*time.Second))
}

// GetUserInfo implements the UserServiceImpl interface.
func GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	return userClient.GetUser(ctx, req)
}
