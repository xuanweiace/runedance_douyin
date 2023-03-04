package db_mysql

import (
	constants "runedance_douyin/pkg/consts"
)

type Relation struct {
	FansID int64 `json:"fans_id,omitempty" gorm:"PRIMARY_KEY; column:fans_id"`
	UserID int64 `json:"user_id,omitempty" gorm:"PRIMARY_KEY; column:user_id"`
}

func (r *Relation) TableName() string {
	return constants.RelationTableName
}
func (r *Relation) PrimaryKey() []string {
	return []string{"fans_id", "user_id"}
}
