package asytask

import (
	"context"
	"encoding/json"
	"errors"
	"runedance_douyin/cmd/message/dal/db_mysql"
	"runedance_douyin/cmd/message/dal/db_redis"
	"runedance_douyin/kitex_gen/message"
	"runedance_douyin/pkg/tools"

	"github.com/hibiken/asynq"
)

// method to handle task
func SyncTaskHandler(ctx context.Context, task *asynq.Task) error {
	var m TaskPlayload
	err := json.Unmarshal(task.Payload(), &m)
	if(err != nil){
		return err
	}
	return TransMsgFromRedisToDB(ctx, m.UserId, m.ToUserId)	
}


// sync DB with redis
func TransMsgFromRedisToDB(ctx context.Context, userId int64, toUserId int64) error{
	redis_msg, err := db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return err
	}
	var messageRecordList []*db_mysql.MessageRecord

	if(len(redis_msg) == 0){								
		return errors.New("no new message chat")
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
			UserToUser: tools.GenerateKeyname(userId, toUserId),
			Content : msg.Content,
			CreateTime : msg.CreateTime, 
		}
		messageRecordList = append(messageRecordList, &msgRecord)
		newLatest = msg.Id
	}
	// update latest sync timestamp 
	db_redis.SetTimestampOfLatestMysql(ctx, userId, toUserId, newLatest)
	return db_mysql.InsertMessage(ctx, messageRecordList)
}