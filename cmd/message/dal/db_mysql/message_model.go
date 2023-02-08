package db_mysql

import (
	"gorm.io/gorm"
	constants "runedance_douyin/pkg/consts"
)


type MessageRecord struct {
	gorm.Model
	Timestamp     int64		   `json:"timestamp"`
	UserToUser    string	   `json:"user_to_user"`
	Content       string	   `json:"content"`
	CreateTime    string	   `json:"create_time"`
}

func (u *MessageRecord) TableName() string {
	return constants.MessageTableName
}


// mod(hash) 分表 