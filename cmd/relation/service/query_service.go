package service

import (
	"context"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/kitex_gen/relation"
	"sync"
)

type QueryService struct {
	ctx context.Context
}

var queryService QueryService
var queryServiceOnce sync.Once

func GetQueryServiceInstance(ctx context.Context) *QueryService {
	queryServiceOnce.Do(func() {
		queryService = QueryService{ctx: ctx}
	})
	return &queryService
}

// todo 确定这里的返回值是relation.User还是db的user
func (q QueryService) GetFollowList(req *relation.GetFollowListRequest) ([]*relation.User, error) {

	// check param
	user_id, err := extract_user_id_from_jwt_token(req.Token)
	if err != nil {
		return nil, err
	}
	// todo 是否需要 去user服务验证user是否存在？
	userids, err := db_mysql.ListFollowidsByUserid(user_id)
	if err != nil {
		return nil, err
	}
	// todo 调用rpc-user去查询用户信息并构造relation.User
	userList := []*relation.User{}
	for _, userid := range userids {
		userList = append(userList, &relation.User{
			Id:            userid,
			Name:          "todo",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      true,
		})
	}
	return userList, nil
}

func (q QueryService) GetFollowerList(req *relation.GetFollowerListRequest) ([]*relation.User, error) {
	// check param
	user_id, err := extract_user_id_from_jwt_token(req.Token)
	if err != nil {
		return nil, err
	}
	// todo 是否需要 去user服务验证user是否存在？
	fansids, err := db_mysql.ListFolloweridsByUserid(user_id)
	if err != nil {
		return nil, err
	}
	// todo 调用rpc-user去查询用户信息并构造relation.User
	fansList := []*relation.User{}
	for _, fansid := range fansids {
		//查看user是否关注了他的fans
		existRelation := q.existRelation(fansid, user_id)
		fansList = append(fansList, &relation.User{
			Id:            fansid,
			Name:          "todo",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      existRelation,
		})
	}
	return fansList, nil
}

func (q QueryService) GetFriendList(req *relation.GetFriendListRequest) ([]*relation.User, error) {
	// check param
	user_id, err := extract_user_id_from_jwt_token(req.Token)
	if err != nil {
		return nil, err
	}
	//先找到user的粉丝
	fansids, err := db_mysql.ListFolloweridsByUserid(user_id)
	followids, err := db_mysql.ListFollowidsByUserid(user_id)
	friendids := intersection_of_id(fansids, followids)
	friendList := []*relation.User{}
	for _, id := range friendids {
		friendList = append(friendList, &relation.User{
			Id:            id,
			Name:          "todo",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      true,
		})
	}
	return friendList, nil
}

// todo 需要绑定在QueryService吗？
func (q QueryService) existRelation(fans_id, user_id int64) bool {
	queryRelation, _ := db_mysql.QueryRelation(fans_id, user_id)
	if queryRelation != nil {
		return true
	}
	return false
}

// 求集合交并补的第三方包 https://github.com/deckarep/golang-set
func intersection_of_id(arr1, arr2 []int64) (ret []int64) {
	sets := make(map[int64]struct{})
	for _, id := range arr1 {
		sets[id] = struct{}{}
	}
	for _, id := range arr2 {
		if _, ok := sets[id]; ok {
			ret = append(ret, id)
		}
	}
	return ret
}
