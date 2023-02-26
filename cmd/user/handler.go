package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runedance_douyin/cmd/user/dal/db_mysql"
	"runedance_douyin/cmd/user/dal/db_redis"
	"runedance_douyin/cmd/user/rpc"
	"runedance_douyin/kitex_gen/user"
	"runedance_douyin/pkg/tools"
	"strconv"
	"sync/atomic"
	"unsafe"
)

const (
	success = 0
	Failed  = 1
)

var userIdSequence int64

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(_ context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	//todo 账户密码校验
	var msg string
	resp := user.NewDouyinUserRegisterResponse()
	//随机生成salt
	salt := tools.RandomStringUtil()
	username := req.Username
	//密码MD5加密
	password := tools.Md5Util(req.Password, salt)
	//更新用户ID
	userIdSequence = db_mysql.GetUserService().FindLastUserId()
	//注册用户
	err := db_mysql.GetUserService().UserRegister(username, password, salt)
	if err != nil {
		//注册失败返回错误信息
		resp.UserId = 0
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//成功注册
		msg = "UserRegister successfully"
		atomic.AddInt64(&userIdSequence, 1)
		token, err := tools.GenToken(username, userIdSequence)
		if err != nil {
			msg = "token generation failed" + err.Error()
			resp.StatusMsg = &msg
			resp.StatusCode = Failed
			return resp, err
		}
		resp.StatusCode = success
		resp.UserId = userIdSequence
		resp.Token = token
		resp.StatusMsg = &msg
	}
	return resp, err
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(_ context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	// todo 账户密码校验
	username := req.Username
	password := req.Password
	//登录验证失败
	//返回：msg:user does not exist | password error
	var msg string
	var resp = user.NewDouyinUserLoginResponse()
	userResp, err := db_mysql.GetUserService().UserLogin(username, password)
	if err != nil {
		resp.StatusCode = Failed
		msg = err.Error()
		resp.StatusMsg = &msg
	} else {
		//登陆成功
		resp.StatusCode = success
		msg = "Successful login"
		resp.StatusMsg = &msg
		resp.UserId = userResp.UserId
		token, _ := tools.GenToken(username, resp.UserId)
		resp.Token = token
	}
	return resp, err
}

type userCache struct {
	UserId        int64  `thrift:"user_id,1,required" frugal:"1,required,i64" json:"user_id"`
	Username      string `thrift:"username,2,required" frugal:"2,required,string" json:"username"`
	FollowCount   *int64 `thrift:"follow_count,3,optional" frugal:"3,optional,i64" json:"follow_count,omitempty"`
	FollowerCount *int64 `thrift:"follower_count,4,optional" frugal:"4,optional,i64" json:"follower_count,omitempty"`
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(_ context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	var msg string
	var err error
	var resp = user.NewDouyinUserResponse()
	resp.User = &user.User{}
	key := strconv.Itoa(int(req.UserId))
	userReply, err := db_redis.RedisGetValue(key)

	var usercache = &userCache{}
	if err == nil {
		json.Unmarshal(S2B(userReply), usercache)
		resp.User.UserId = usercache.UserId
		resp.User.Username = usercache.Username
		resp.User.FollowerCount = usercache.FollowerCount
		resp.User.FollowCount = usercache.FollowCount
		resp.User.IsFollow, _ = rpc.ExistRelation(req.MyUserId, req.UserId)
	} else {
		resp.User, err = db_mysql.GetUserService().GetUserById(req.UserId, req.MyUserId)
		go func() {
			usercache.UserId = resp.User.UserId
			usercache.Username = resp.User.Username
			usercache.FollowCount = resp.User.FollowCount
			usercache.FollowerCount = resp.User.FollowerCount
			value, err := json.Marshal(usercache)
			if err == nil {
				db_redis.RedisCacheString(key, B2S(value), db_redis.DefaultExpirationTime)
			}
		}()

	}
	if err == nil {
		resp.StatusCode = success
		msg = "GetUser success"
		resp.StatusMsg = &msg
		return resp, nil
	} else {
		msg = "get user information failed"
	}
	if err != nil {
		resp.StatusCode = Failed
	}
	resp.StatusMsg = &msg
	return resp, err
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(_ context.Context, req *user.DouyinUserUpdateRequest) (*user.DouyinUserUpdateResponse, error) {
	resp := user.NewDouyinUserUpdateResponse()
	var err error
	fmt.Println("req.Fo:", req.Followdiff, req.Followerdiff)

	// 删除redis缓存
	key := strconv.Itoa(int(req.UserId))
	go db_redis.RedisDo("del", key)

	if req.Followdiff != 0 {
		err = db_mysql.GetUserService().UpdateUserFollow(req.UserId, req.Followdiff)
		if err != nil {
			resp.StatusCode = Failed
			msg := "UpdateUserFollow failed"
			resp.StatusMsg = &msg
			return resp, err
		}
	}
	if req.Followerdiff != 0 {
		err = db_mysql.GetUserService().UpdateUserFollower(req.UserId, req.Followerdiff)
		if err != nil {
			resp.StatusCode = Failed
			msg := "UpdateUserFollower failed"
			resp.StatusMsg = &msg
			return resp, err
		}
	}
	resp.StatusCode = success
	msg := "update done"
	resp.StatusMsg = &msg
	return resp, nil
}
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func S2B(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
