package db_mysql

import (
	"errors"
	"fmt"
	"runedance_douyin/pkg/tools"
	"sync"
)

var (
	userService     UserService
	userServiceOnce sync.Once
)

type UserService interface {
	UserLogin(username string, password string) (*User, error)
	Register(username string, password string, userId int64, salt string) error
	FindLastUserId() int64
}
type UserServiceImpl struct {
	userDao UserDao
}

func GetUserService() UserService {
	userServiceOnce.Do(func() {
		userService = &UserServiceImpl{
			userDao: GetUserDao(),
		}
	})
	return userService
}

func (u *UserServiceImpl) UserLogin(username string, password string) (*User, error) {
	user, err := u.userDao.FindByName(username)
	if err != nil {
		return nil, errors.New("user does not exist")
	}
	//校验密码
	fmt.Println(user)
	//MD5加密验证
	password = tools.Md5Util(password, user.Salt)
	if password != user.Password {
		return nil, errors.New("password error")
	}
	return user, nil
}

// 注册用户
// 1.先判断表里有没有用户 如果有就提示用户存在
// 2.判断用户名是否违法或者合规（暂未实现）
// 3.注册用户

func (u *UserServiceImpl) Register(username string, password string, userId int64, salt string) error {
	//判断用户是否已经注册
	_, err := u.userDao.FindByName(username)
	if err == nil {
		return errors.New("user does not exist")
	}
	//添加用户
	user := User{
		UserId:        userId,
		Username:      username,
		Password:      password,
		Salt:          salt,
		FollowCount:   0,
		FollowerCount: 0,
	}
	e := u.userDao.AddUser(&user)
	if e != nil {
		return errors.New("user regist failed")
	}
	return nil
}

// 返回当前最大的用户ID

func (u *UserServiceImpl) FindLastUserId() int64 {
	return u.userDao.LastId()
}
