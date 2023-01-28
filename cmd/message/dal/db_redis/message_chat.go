package db_redis

import (
	"context"
)

func GetMessageChatJson(ctx context.Context, userId string, toUserId string) ([]string, error){
	keyname1 := GenerateKeyname(userId, toUserId)
	// get messageRecord json info by keyname
	jsonList, err:= Rdb.LRange(ctx, keyname1, 0, Rdb.LLen(ctx, keyname1).Val()).Result()
	return jsonList, err
}

