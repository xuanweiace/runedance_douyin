package db_redis

import (
	"context"
	"encoding/json"
	"runedance_douyin/kitex_gen/message"
	"runedance_douyin/pkg/tools"
)

// get messageRecord json info by keyname in redis
func GetMessageChatJson(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname := tools.GenerateKeyname(userId, toUserId)
	jsonList, err:= Rdb.LRange(ctx, keyname, 0, Rdb.LLen(ctx, keyname).Val()).Result()
	return jsonList, err
	
	
	// log.Printf(keyname)
	
}


// insert message chat into redis
func LoadMessageChat(ctx context.Context, userId int64, toUserId int64, msgList []*message.Message) error{
	keyname := tools.GenerateKeyname(userId, toUserId)
	var Err error
	for _, val := range msgList {
		// encode message into json
		jsonStr, err := json.Marshal(val)
		if(err != nil){
			Err = err
			continue
		}
		Err = Rdb.LPush(ctx, keyname, string(jsonStr)).Err()
		Rdb.Expire(ctx, keyname, 3600)     	// refresh keyname expiration 
	}
	return Err
}



// key userId - toUserid : [taskIda, ]
// a. Taska (t1 + 5min)  b. taskb (t2 + 5min)


