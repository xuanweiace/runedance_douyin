package db_redis

import (
	"context"
	"encoding/json"
	"runedance_douyin/cmd/message/dal/db_mysql"
	"runedance_douyin/pkg/tools"
	"time"
)

// get messageRecord json info by keyname in redis
func GetMessageChatJson(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname := tools.GenerateKeyname(userId, toUserId)
	jsonList, err:= RdbCluster.LRange(ctx, keyname, 0, RdbCluster.LLen(ctx, keyname).Val()).Result()
	return jsonList, err
	// log.Printf(keyname)
}


// insert message chat into redis
func LoadMessageChat(ctx context.Context, userId int64, toUserId int64, msgList []*db_mysql.MessageRecord) error{
	keyname := tools.GenerateKeyname(userId, toUserId)
	var Err error
	for _, val := range msgList {
		// encode message into json
		jsonStr, err := json.Marshal(val)
		if(err != nil){
			Err = err
			continue
		}

		// add to the head of list
		Err = RdbCluster.LPush(ctx, keyname, string(jsonStr)).Err()
		RdbCluster.Expire(ctx, keyname, time.Hour)     	// refresh keyname expiration 
	}
	return Err
}



// key userId - toUserid : [taskIda, ]
// a. Taska (t1 + 5min)  b. taskb (t2 + 5min)


