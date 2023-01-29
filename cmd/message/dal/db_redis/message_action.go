package db_redis

import(
	"context"
	"time"
	"encoding/json"
	"runedance_douyin/kitex_gen/message"
)


func HandleMessageSend (ctx context.Context, userId int64, toUserId int64, actionType int32, content string)  error{
	keyname := GenerateKeyname(userId, toUserId)
	cur_time := time.Now().String()
	// store message into a map
	m := message.Message{
		Id: time.Now().Unix(),
		Content: content,
		CreateTime: &cur_time,
	}
	// convert m to json
	jsonStr, err := json.Marshal(m)
	if(err != nil){
		return err
	}
	// store json string in redis with key being the keyname generated based on userId and toUserId
	error := Rdb.LPush(ctx, keyname, string(jsonStr)).Err()
	return error
}

