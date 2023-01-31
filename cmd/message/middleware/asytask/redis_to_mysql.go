package asytask

import (
	"runedance_douyin/cmd/message/dal/db_redis"
	"runedance_douyin/cmd/message/dal/db_mysql"
	"runedance_douyin/kitex_gen/message"
	"context"
	"encoding/json"
	"runedance_douyin/pkg/tools"
)


// sync DB with redis
func TransMsgFromRedisToDB(ctx context.Context, userId int64, toUserId int64) error{
	redis_msg, err := db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return err
	}
	var messageRecordList []*db_mysql.MessageRecord

	if(redis_msg == nil){								
		return nil
	}

	// get latest sync timestamp 
	latestTime, err := db_redis.GetTimestampOfLatestMysql(ctx, userId, toUserId)

	if(err != nil){
		return err
	}
	var newLatest int64
	for _, val := range redis_msg {
		var msg message.Message
		err := json.Unmarshal([]byte(val), &msg)			// decode json into Message struct
		if(err != nil){
			continue
		}
		if(msg.Id < latestTime){
			continue
		}
		msgRecord := db_mysql.MessageRecord{
			ID : msg.Id,
			UserToUser: tools.GenerateKeyname(userId, toUserId),
			Content : msg.Content,
			CreateTime : *msg.CreateTime, 
		}
		messageRecordList = append(messageRecordList, &msgRecord)
		newLatest = msg.Id
	}
	// update latest sync timestamp 
	db_redis.SetTimestampOfLatestMysql(ctx, userId, toUserId, newLatest)
	return db_mysql.InsertMessage(ctx, messageRecordList)
}