package db_mysql

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"runedance_douyin/pkg"
	constants "runedance_douyin/pkg/consts"
	"sync"
)

func MySQLInit() {
	pkg.InitDB(constants.MySQLDefaultDSN)
}

var (
	userDao     UserDao
	userDaoOnce sync.Once
)

type UserDao interface {
	AddUser(user *User) error
	FindByName(name string) (*User, error)
	FindById(userid int64) (*User, error)
	LastId() int64
	UpdateFollow(userid int64, followDiff int64) error
	UpdateFollower(userid int64, followerDiff int64) error
}
type UserDaoImpl struct {
	db  *gorm.DB
	rec redis.Conn
}

func GetUserDao() UserDao {
	userDaoOnce.Do(func() {
		userDao = &UserDaoImpl{
			db:  pkg.GetDB(),
			rec: pkg.GetRec(),
		}
	})
	return userDao
}

// AddUser 添加用户
// 参数 user User结构体指针
func (u *UserDaoImpl) AddUser(user *User) error {
	if err := u.db.Model(&User{}).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindByName 根据用户名查找用户
// 参数 name string类型 用户名
func (u *UserDaoImpl) FindByName(name string) (*User, error) {
	var user User
	if err := u.db.Model(&User{}).Where("name = ?", name).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

// FindById 根据用户id查找用户
// 参数 userid int64类型 用户id
func (u *UserDaoImpl) FindById(userid int64) (*User, error) {
	var user User
	if err := u.db.Model(&User{}).Where("user_id = ?", userid).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

// 通过主键查询最后一条记录
// 返回当前表内的最大ID

func (u *UserDaoImpl) LastId() int64 {
	var user User
	if err := u.db.Model(&User{}).Last(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//表内没有数据默认为1
		return 1
	}
	return user.UserId
}

func (u *UserDaoImpl) UpdateFollow(userid int64, followDiff int64) error {
	var getUser = &User{}
	err := u.db.Model(&User{}).Where("user_id=?", userid).First(getUser)
	if err != nil {
		return err.Error
	}
	return u.db.Model(&User{}).Where("user_id=?", userid).Update("follow_count=?", followDiff+getUser.FollowCount).Error
}

func (u *UserDaoImpl) UpdateFollower(userid int64, followerDiff int64) error {
	var getUser = &User{}
	err := u.db.Model(&User{}).Where("user_id=?", userid).First(getUser)
	if err != nil {
		return err.Error
	}
	return u.db.Model(&User{}).Where("user_id=?", userid).Update("follower_count=?", followerDiff+getUser.FollowerCount).Error
}
