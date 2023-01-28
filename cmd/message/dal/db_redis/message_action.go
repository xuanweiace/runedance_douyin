package db_redis

import(
	"context"
	"time"
	"encoding/json"
)

type UserMessageRecord struct {
	UserId		int64	 `redis:"userId"`
	ToUserId    int64    `redis:"toUserId"`
	Content     string   `redis:"content"`
	CreateTime  string   `redis:"createTime"`
}


func HandleMessageSend (ctx context.Context, userId string, toUserId string, actionType int32, content string)  error{
	keyname := GenerateKeyname(userId, toUserId)
	
	// store message into a map
	m := make(map[string]string)
	m["userId"] = userId
	m["toUserId"] = toUserId
	m["content"] = content
	m["createTime"] = time.Now().String()

	// convert map to json
	jsonStr, err := json.Marshal(m)
	if(err != nil){
		return err
	}
	// store json string in redis with key being the keyname generated based on userId and toUserId
	error := Rdb.LPush(ctx, keyname, string(jsonStr)).Err()
	return error
}

// generation keyname by user_id and to_user_id
func GenerateKeyname(userId string, toUserId string) (string){
	str := userId + "-" + toUserId
	return str
}