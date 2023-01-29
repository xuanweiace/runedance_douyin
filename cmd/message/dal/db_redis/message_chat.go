package db_redis

import (
	"context"
	"runedance_douyin/kitex_gen/message"
	"runedance_douyin/pkg/tools"
	"encoding/json"
)

func GetMessageChatJson(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname := tools.GenerateKeyname(userId, toUserId)
	// check if keyname exist in redis
	exist := Rdb.Exists(ctx, keyname).Val()

	if(exist == 2){				// get messageRecord json info by keyname in redis
		jsonList, err:= Rdb.LRange(ctx, keyname, 0, Rdb.LLen(ctx, keyname).Val()).Result()

		// insert into mysqlDB and release redis memory

		return jsonList, err
	}
	return nil, nil    // keyname does not exist or stores nothing
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
	}
	return Err

}





