package mysql

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"runedance_douyin/pkg"
	"sync"
)

var (
	favoriteDao     FavoriteDao
	favoriteDaoOnce sync.Once
)

type FavoriteDao interface {
	FindFavorite(uid int64, vid int64) (*Favorite, error)
	AddFavorite(favorite *Favorite) error
	UpdateFavorite(id string, action int32) error
	GetFavoriteList(uid int64) ([]int64, error)
}

type FavoriteDaoImpl struct {
	db  *gorm.DB
	rec redis.Conn
}

func GetFavoriteDao() FavoriteDao {
	favoriteDaoOnce.Do(
		func() {
			favoriteDao = &FavoriteDaoImpl{
				db:  pkg.GetDB(),
				rec: pkg.GetRec(),
			}
		})
	return favoriteDao
}

// 1 FavoriteAction
// 根据uid及vid查找记录
func (f *FavoriteDaoImpl) FindFavorite(uid int64, vid int64) (*Favorite, error) {
	var favorite Favorite
	if err := f.db.Table("favorite").Where("uid=?", uid).Where("vid=?", vid).First(&favorite).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &favorite, nil
}

// 插入记录
func (f *FavoriteDaoImpl) AddFavorite(favorite *Favorite) error {
	if err := f.db.Table("favorite").Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

// 修改记录
func (f *FavoriteDaoImpl) UpdateFavorite(id string, action int32) error {
	f.db.Table("favorite").Where("id=?", id).Update("action", action)
	return nil
}

// 2 GetFavoriteList
// 根据uid及action查找记录 vid list
func (f *FavoriteDaoImpl) GetFavoriteList(uid int64) ([]int64, error) {
	var res []int64
	var record []Favorite
	f.db.Table("favorite").Where("uid=?", uid).Where("action=?", 1).Find(&record)
	for _, cur := range record {
		res = append(res, cur.Vid)
	}
	return res, nil
}
