package service

import (
	"context"
	"errors"
	"runedance_douyin/cmd/relation/dal/db_redis"
	"runedance_douyin/cmd/relation/rpc"
	"runedance_douyin/kitex_gen/relation"
	"runedance_douyin/kitex_gen/user"
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
	user_id := req.UserId

	// 去user服务验证user是否存在
	if _, err := rpc.GetUser(user_id); err != nil {
		return nil, err
	}
	userids, err := db_redis.ListFollowidsByUserid(q.ctx, user_id)
	// userids, err := db_mysql.ListFollowidsByUserid(user_id)
	if err != nil {
		return nil, err
	}
	//调用rpc-user去查询用户信息并构造relation.User
	userList := []*relation.User{}
	for _, userid := range userids {
		usr := fill_user_info(userid)
		usr.IsFollow = true
		userList = append(userList, usr)
	}
	return userList, nil
}

func (q QueryService) GetFollowerList(req *relation.GetFollowerListRequest) ([]*relation.User, error) {
	// check param
	user_id := req.UserId

	if user_id != req.UserId {
		return nil, errors.New("param UserId error")
	}

	// 去user服务验证user是否存在
	if _, err := rpc.GetUser(user_id); err != nil {
		return nil, err
	}
	fansids, err := db_redis.ListFolloweridsByUserid(q.ctx, user_id)
	// fansids, err := db_mysql.ListFolloweridsByUserid(user_id)
	if err != nil {
		return nil, err
	}
	// 调用rpc-user去查询用户信息并构造relation.User
	fansList := []*relation.User{}
	for _, fansid := range fansids {
		fans := fill_user_info(fansid)
		//查看user是否关注了他的fans
		existRelation := q.existRelation(user_id, fansid)
		fans.IsFollow = existRelation
		fansList = append(fansList, fans)
	}
	return fansList, nil
}

func (q QueryService) GetFriendList(req *relation.GetFriendListRequest) ([]*relation.User, error) {
	// check param
	user_id := req.UserId

	// 去user服务验证user是否存在
	if _, err := rpc.GetUser(user_id); err != nil {
		return nil, err
	}

	//先找到user的粉丝
	fansids, err := db_redis.ListFolloweridsByUserid(q.ctx, user_id)
	if err != nil {
		return nil, err
	}
	followids, err := db_redis.ListFollowidsByUserid(q.ctx, user_id)
	if err != nil {
		return nil, err
	}
	friendids := intersection_of_id(fansids, followids)
	friendList := []*relation.User{}
	for _, id := range friendids {
		friendList = append(friendList, fill_user_info(id))
	}
	return friendList, nil
}

func (q QueryService) ExistRelation(req *relation.ExistRelationRequest) (bool, error) {
	b := q.existRelation(req.FromUserId, req.ToUserId)
	return b, nil
}

// todo 需要绑定在QueryService吗？
func (q QueryService) existRelation(fans_id, user_id int64) bool {
	queryRelation, _ := db_redis.QueryRelation(q.ctx, fans_id, user_id)
	// queryRelation, _ := db_mysql.QueryRelation(fans_id, user_id)
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

func fill_user_info(id int64) *relation.User {
	usr, err := rpc.GetUser(id)
	if err != nil {
		usr = &user.User{
			UserId:        id,
			Username:      "user not exist",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		}
	}
	return &relation.User{
		Id:            id,
		Name:          usr.Username,
		FollowCount:   usr.FollowCount,
		FollowerCount: usr.FollowerCount,
		IsFollow:      true,
	}
}
