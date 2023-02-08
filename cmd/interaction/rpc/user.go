package rpc

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"time"
)

var userClient userservice.Client

func initUser() {
	// 好像0.0.0.0不行？在mac上
	c, err := userservice.NewClient(constants.UserServiceName, client.WithHostPorts("127.0.0.1:8888"))

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
