package db_mysql

import (
	"github.com/go-sql-driver/mysql"
	"log"
)

func CreateRelation(relation *Relation) error {
	err := db.Create(relation).Error
	//错误修复，如果是Duplicate entry的err则放过
	//todo 更好的处理方式？
	if x, ok := err.(*mysql.MySQLError); ok {
		if x.Number == 0x426 {
			err = nil
		}
	}
	log.Printf("[db_mysql.CreateRelation] relation=%+v, err=%#v", relation, err)
	return err
}

func DeleteRelation(relation *Relation) error {
	//err := db.Delete(relation).Error // 联合主键不能这么删除？
	err := db.Where("fans_id = ? and user_id = ?", relation.FansID, relation.UserID).Delete(relation).Error
	log.Printf("[db_mysql.DeleteRelation] relation=%+v, err=%#v", relation, err)
	return err
}
