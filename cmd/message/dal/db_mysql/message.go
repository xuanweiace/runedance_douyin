package db_mysql

import (
	"context"
	"runedance_douyin/pkg/tools"
)


func InsertMessage(ctx context.Context, msgResordList []*MessageRecord) error {
	err := db.WithContext(ctx).Create(msgResordList).Error
	if(err != nil){
		return err
	}
	return nil
}

func GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*MessageRecord, error) {
	var result []*MessageRecord
	keyname := tools.GenerateKeyname(userId, toUserId)
	rows, err := db.WithContext(ctx).Model(&MessageRecord{}).Where("user_to_user = ?", keyname).Rows()
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


