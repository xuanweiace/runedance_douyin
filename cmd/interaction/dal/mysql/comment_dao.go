package mysql

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"runedance_douyin/pkg"
	"sync"
)

var (
	commentDao     CommentDao
	commentDaoOnce sync.Once
)

type CommentDao interface {
	AddComment(comment *Comment) error
	DeleteComment(cid int64) error
	FindCommentListByVid(vid int64) ([]Comment, error)
	FindComment(uid int64, vid int64) (*Comment, error)
}

type CommentDaoImpl struct {
	db  *gorm.DB
	rec redis.Conn
}

func GetCommentDao() CommentDao {
	commentDaoOnce.Do(
		func() {
			commentDao = &CommentDaoImpl{
				db:  pkg.GetDB(),
				rec: pkg.GetRec(),
			}
		})
	return commentDao
}

// 3 CommentAction
// 插入评论
func (c *CommentDaoImpl) AddComment(comment *Comment) error {
	if err := c.db.Table("comment").Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// 修改记录
func (c *CommentDaoImpl) DeleteComment(cid int64) error {
	c.db.Table("comment").Where("Id=?", cid).Delete(&Comment{})
	return nil
}

// 4 GetCommentList
// 根据vid查找记录 list 发布时间倒序 大的在后
func (c *CommentDaoImpl) FindCommentListByVid(vid int64) ([]Comment, error) {
	var res []Comment
	c.db.Table("comment").Order("Content_date desc").Where("vid=?", vid).Find(&res)
	return res, nil
}

// 根据uid vid查找评论
func (c *CommentDaoImpl) FindComment(uid int64, vid int64) (*Comment, error) {
	var comment Comment
	if err := c.db.Table("comment").Where("uid=?", uid).Where("vid=?", vid).First(&comment).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &comment, nil
}
