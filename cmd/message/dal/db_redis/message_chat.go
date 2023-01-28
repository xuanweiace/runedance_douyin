package db_redis

import(
	"context"
	"runedance_douyin/kitex_gen/message"
	"encoding/json"
	"time"
)

func GetMessageChat(ctx context.Context, userId string, toUserId string) ([]*message.Message, error){
	var result []*message.Message
	keyname := GenerateKeyname(userId, toUserId)
	// get message record json string by keyname
	recordList, err:= Rdb.LRange(ctx, keyname, 0, Rdb.LLen(ctx, keyname).Val()).Result()
	if(err != nil){
		return result, err
	}
	// decode json into UserMessageRecord struct
	for _, val := range recordList {
		var record UserMessageRecord
		var msg message.Message
		err := json.Unmarshal([]byte(val), &record)
		if(err != nil){
			continue
		}
		// set msg fields
		msg.Id = time.Now().Unix()
		msg.Content = record.Content
		msg.CreateTime = &record.CreateTime
		result = append(result, &msg)
	}
	return result, nil
}

