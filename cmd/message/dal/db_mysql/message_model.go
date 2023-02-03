package db_mysql

import (
	"gorm.io/gorm"
)

const TableNameMessage string = "message"

type MessageRecord struct {
	gorm.Model
	Timestamp     int64		   `json:"timestamp"`
	UserToUser    string	   `json:"user_to_user"`
	Content       string	   `json:"content"`
	CreateTime    string	   `json:"create_time"`
}

func (u *MessageRecord) TableName() string {
	return TableNameMessage
}


// mod(hash) 分表 