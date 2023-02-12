package db_redis

import(
	"context"
	"time"
	"encoding/json"
	"runedance_douyin/kitex_gen/message"
	"runedance_douyin/pkg/tools"
	"strconv"
)

func HandleMessageSend (ctx context.Context, userId int64, toUserId int64, actionType int32, content string)  error{
	keyname := tools.GenerateKeyname(userId, toUserId)
	// store message into a map
	curTime := time.Now()
	m := message.Message{
		Id: curTime.Unix(),
		Content: content,
		CreateTime: curTime.String(),
	}

	// encode m to json
	jsonStr, err := json.Marshal(m)
	if(err != nil){
		return err
	}
	
	// store json string in redis with key being the keyname generated based on userId and toUserId
	// add to the tail of list
	error := Rdb.RPush(ctx, keyname, string(jsonStr)).Err()
	Rdb.ExpireAt(ctx, keyname, curTime.Add(time.Hour))       		// refreshing exprie time to 1h
	return error
}


// get latest sync timestamp 
func GetTimestampOfLatestMysql(ctx context.Context, userId int64, toUserId int64) (int64, error){
	key := tools.GenerateKeyname(userId, toUserId) + "-latest"

	if(Rdb.Exists(ctx, key).Val() == 2){
		str := Rdb.Get(ctx, key).String()
		return strconv.ParseInt(str, 10, 64)
	}
	return -1, nil
}

// set
func SetTimestampOfLatestMysql(ctx context.Context, userId int64, toUserId int64, timestamp int64) error {
	key := tools.GenerateKeyname(userId, toUserId) + "-latest"
	return Rdb.Set(ctx, key, strconv.FormatInt(timestamp, 10), time.Hour).Err()
}




