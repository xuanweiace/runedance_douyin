package db_redis

import (
	"context"
	"strconv"
)

func GetMessageChatJson(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname1 := GenerateKeyname(userId, toUserId)
	// get messageRecord json info by keyname
	jsonList, err:= Rdb.LRange(ctx, keyname1, 0, Rdb.LLen(ctx, keyname1).Val()).Result()
	return jsonList, err
}


// generation keyname by user_id and to_user_id
func GenerateKeyname(userId int64, toUserId int64) (string){
	var str string
	if(userId < toUserId){
		str = strconv.FormatInt(userId, 10) + "-" + strconv.FormatInt(toUserId, 10)
		return str
	}
	str = strconv.FormatInt(toUserId, 10) + "-" + strconv.FormatInt(userId, 10)
	return str
}