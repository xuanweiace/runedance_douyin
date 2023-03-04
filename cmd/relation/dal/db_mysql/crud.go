package db_mysql

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateRelation(tx *gorm.DB, relation *Relation) error {
	if tx == nil {
		tx = db
	}
	err := tx.Create(relation).Error
	//不进行错误修复，如果是Duplicate entry的err也需要抛出来
	// if x, ok := err.(*mysql.MySQLError); ok {
	// 	if x.Number == 0x426 {
	// 		err = nil
	// 	}
	// }
	log.Printf("[db_mysql.CreateRelation] relation=%+v, err=%#v", relation, err)
	return err
}
func BatchCreate(relation []*Relation) error {
	return db.Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(relation, len(relation)).Error
}
func BatchDelete(relation []*Relation) error {
	// primaryKeys := make([][]interface{}, 0)
	// for _, rel := range relation {
	// 	primaryKeys = append(primaryKeys, []interface{}{rel.FansID, rel.UserID})
	// }
	if len(relation) == 0 {
		return nil
	}

	sql := "delete from relation where (fans_id, user_id) in ((?,?)"
	li := []interface{}{}
	for i, rel := range relation {
		if i > 0 {
			sql += ",(?,?)"
		}
		li = append(li, rel.FansID, rel.UserID)
	}
	sql += ")"
	fmt.Println("sql:", sql)
	fmt.Println("li:", li)
	return db.Exec(sql, li...).Error
	// primaryKeys := make([]map[string]interface{}, 0)
	// for _, rel := range relation {
	// 	primaryKeys = append(primaryKeys, map[string]interface{}{"fans_id": rel.FansID, "user_id": rel.UserID})
	// }
	// return db.Where(primaryKeys).Delete(&Relation{}).Error
}
func DeleteRelation(tx *gorm.DB, relation *Relation) error {
	if tx == nil {
		tx = db
	}
	//err := db.Delete(relation).Error // 联合主键不能这么删除？
	err := tx.Where("fans_id = ? and user_id = ?", relation.FansID, relation.UserID).Delete(relation).Error
	log.Printf("[db_mysql.DeleteRelation] relation=%+v, err=%#v", relation, err)
	return err
}

// 获取一个用户的关注id列表
func ListFollowidsByUserid(userid int64) ([]int64, error) {
	follows := make([]*Relation, 0)
	if err := db.Where("fans_id = ?", userid).Find(&follows).Error; err != nil {
		return nil, err
	}
	res := make([]int64, 0)
	for _, v := range follows {
		res = append(res, v.UserID)
	}
	return res, nil
}

// 获取一个用户的粉丝id列表
func ListFolloweridsByUserid(userid int64) ([]int64, error) {
	followers := make([]*Relation, 0)
	if err := db.Where("user_id = ?", userid).Find(&followers).Error; err != nil {
		return nil, err
	}
	res := make([]int64, 0)
	for _, v := range followers {
		res = append(res, v.FansID)
	}
	return res, nil
}

// 查看是否存在一对relation
// 如果不存在则err=record not found（如果用Find方法则找不到也不会有err，返回零值）
func QueryRelation(fansid, userid int64) (*Relation, error) {
	relations := &Relation{}
	//不加&会panic
	if err := db.Where("fans_id = ? and user_id = ?", fansid, userid).First(relations).Error; err == nil {
		return relations, err
	} else {
		return nil, err
	}

}

func ExecFuncInTransaction(f func(tx *gorm.DB) error) error {
	// err := f
	err := db.Transaction(f)
	return err
}
