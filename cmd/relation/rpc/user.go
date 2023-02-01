package rpc

import (
	"errors"
	"github.com/cloudwego/kitex/client"
	"gorm.io/gorm"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/kitex_gen/user/userservice"
	constants "runedance_douyin/pkg/consts"
)

var userClient userservice.Client

func initUser() {
	c, err := userservice.NewClient(constants.UserServiceName, client.WithHostPorts("0.0.0.0:8888"))

	if err != nil {
		panic(err)
	}
	userClient = c
}

func GetUser(user_id int64) (*user.User, error) {
	usr := user.User{}
	if err := db_mysql.DB.Table("user").Where("user_id = ?", user_id).First(&usr).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &usr, nil
	//token, _ := tools.GenToken("todo", user_id)
	//
	//request := user.DouyinUserRequest{
	//	UserId: user_id,
	//	Token:  token,
	//}
	//response, err := userClient.GetUser(context.Background(), &request, callopt.WithRPCTimeout(10*time.Second))
	//if err != nil {
	//	return nil, err
	//}
	//if response.StatusCode != errnos.CodeSuccess {
	//	return nil, errors.New(*response.StatusMsg)
	//}
	//return response.GetUser(), nil
}

type DBUser struct {
	UserId        int64
	Username      string
	Password      string
	Avatar        string
	Salt          string
	FollowCount   int64
	FollowerCount int64
}

func (u *DBUser) TableName() string {
	return "user"
}
