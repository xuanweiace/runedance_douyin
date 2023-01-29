package db_mysql

import (
	"gorm.io/gorm"
)

const TableNameMessage string = "message"

type MessageRecord struct {
	gorm.Model
	ID      	  int64        `json:"id"`
	UserToUser    string	   `json:"user_to_user"`
	Content       string	   `json:"content"`
	CreateTime    string	   `json:"create_time"`
}

func (u *MessageRecord) TableName() string {
	return TableNameMessage
}