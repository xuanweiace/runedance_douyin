package db_redis

import (
	"context"
	"encoding/json"
	"runedance_douyin/pkg/tools"
	"strconv"
	"time"
	"runedance_douyin/cmd/message/dal/db_mysql"
)


func HandleMessageSend (ctx context.Context, userId int64, toUserId int64, actionType int32, content string)  error{
	keyname := tools.GenerateKeyname(userId, toUserId)
	// store message into a map
	curTime := time.Now()
	m := db_mysql.MessageRecord{
		Timestamp: curTime.Unix(),
		Sender: userId,
		Content: content,
	}

	// encode m to json
	jsonStr, err := json.Marshal(m)
	if(err != nil){
		return err
	}
	
	// store json string in redis with key being the keyname generated based on userId and toUserId
	// add to the tail of list
	error := RdbCluster.RPush(ctx, keyname, string(jsonStr)).Err()
	RdbCluster.ExpireAt(ctx, keyname, curTime.Add(time.Hour))       		// refreshing exprie time to 1h
	return error
}


// get latest sync timestamp 
func GetTimestampOfLatestMysql(ctx context.Context, userId int64, toUserId int64) (int64, error){
	key := tools.GenerateKeyname(userId, toUserId) + "-latest"

	if(RdbCluster.Exists(ctx, key).Val() == 2){
		str := RdbCluster.Get(ctx, key).String()
		return strconv.ParseInt(str, 10, 64)
	}
	return -1, nil
}

// set
func SetTimestampOfLatestMysql(ctx context.Context, userId int64, toUserId int64, timestamp int64) error {
	key := tools.GenerateKeyname(userId, toUserId) + "-latest"
	return RdbCluster.Set(ctx, key, strconv.FormatInt(timestamp, 10), time.Hour).Err()
}




