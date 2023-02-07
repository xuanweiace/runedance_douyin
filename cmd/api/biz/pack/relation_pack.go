package pack

import (
	"runedance_douyin/cmd/api/biz/model/douyin"
	"runedance_douyin/kitex_gen/relation"
)

//用于将rpc得到的结果打包成api的response。做类型转换和err封装，避免handler里有大量重复代码

func ConvertUserlist(userList []*relation.User) (ul []*douyin.User) {
	ul = make([]*douyin.User, 0)
	for _, user := range userList {
		ul = append(ul, &douyin.User{
			ID:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		})
	}
	return
}
