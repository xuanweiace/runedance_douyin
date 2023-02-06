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
)

var userClient userservice.Client

func initUser() {
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
