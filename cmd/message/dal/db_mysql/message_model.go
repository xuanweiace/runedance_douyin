package db_mysql

import (
	"gorm.io/gorm"
	constants "runedance_douyin/pkg/consts"
)

type MessageRecord struct {
	gorm.Model
	Timestamp     int64		   `json:"timestamp"`
	UserToUser    string	   `json:"user_to_user"`
	Sender 		  int64	       `json:"sender"`
	Content       string	   `json:"content"`
}

func (u *MessageRecord) TableName() string {
	return constants.MessageTableName
}

