package rpc

import (
	"context"
	"errors"
	"fmt"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
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
	fmt.Println("user 这个 rpc 服务 启动成功", c)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func GetUser(user_id int64) (*user.User, error) {
	// return &user.User{UserId: user_id, Username: "test_name"}, nil
	request := user.DouyinUserRequest{
		UserId:   user_id,
		MyUserId: user_id,
	}
	response, err := userClient.GetUser(context.Background(), &request, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		return nil, err
	}
	if response.StatusCode != errnos.CodeSuccess {
		return nil, errors.New(*response.StatusMsg)
	}
	return response.GetUser(), nil
}

func UpdateUser(user_id, follow_diff, follower_diff int64) (bool, error) {
	// return true, nil
	request := user.DouyinUserUpdateRequest{
		UserId:       user_id,
		Followdiff:   follow_diff,
		Followerdiff: follower_diff,
	}
	response, err := userClient.UpdateUser(context.Background(), &request, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		return false, err
	}
	if response.StatusCode != errnos.CodeSuccess {
		err = fmt.Errorf("[Update User] update failed, StatusCode=%d", response.StatusCode)
		return false, err
	}
	return true, nil
}
