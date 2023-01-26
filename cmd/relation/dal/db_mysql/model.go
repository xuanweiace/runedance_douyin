package db_mysql

import (
	constants "runedance_douyin/pkg/consts"
)

type Relation struct {
	FansID int64 `json:"fans_id,omitempty" gorm:"column:fans_id"`
	UserID int64 `json:"user_id,omitempty" gorm:"column:user_id"`
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}
