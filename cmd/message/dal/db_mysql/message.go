package db_mysql

import (
	"context"
	"runedance_douyin/pkg/tools"
)


func InsertMessage(ctx context.Context, msgResordList []*MessageRecord, userId int64, toUserId int64) error {
	// keyname := tools.GenerateKeyname(userId, toUserId)
	// var err error
	// for _, odr:= range msgResordList {
	// 	err = db.Create(odr).Error
	// }
	err := db.WithContext(ctx).Model(&MessageRecord{}).Create(msgResordList).Error
	if(err != nil){
		return err
	}
	return err
}


func GetMessageChat(ctx context.Context, userId int64, toUserId int64, limit int, earlistTime int64) ([]*MessageRecord, error) {
	var result []*MessageRecord
	keyname := tools.GenerateKeyname(userId, toUserId)
	rows, err := db.WithContext(ctx).Model(&MessageRecord{}).Where("user_to_user = ?", keyname).Where("timestamp < ?", earlistTime).Order("timestamp desc").Limit(limit).Rows()
	if(err != nil){
		return result, err
	}
	
	defer rows.Close()
	// scan rows to MessageRecord struct
	for rows.Next() {
		var record MessageRecord
		db.ScanRows(rows, &record)
		result = append(result, &record)
	}
	return result, nil
}


