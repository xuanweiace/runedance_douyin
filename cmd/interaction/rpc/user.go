package rpc

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"time"
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

func GetUser(user_id int64) (*user.User, error) {

	request := user.DouyinUserRequest{
		UserId:   user_id,
		MyUserId: user_id,
	}
	response, err := userClient.GetUser(context.Background(), &request, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != errnos.CodeSuccess {
		return nil, errors.New(*response.StatusMsg)
	}
	return response.GetUser(), nil
}
