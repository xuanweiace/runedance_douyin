package db_redis

import(
	"context"
	"runedance_douyin/pkg/tools"
)

func GetPendingTaskIDs(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	// check if keyname exist in redis
	exist := Rdb.Exists(ctx, keyname).Val()
	if(exist == 2){				// get pending task id json info by keyname in redis
		jsonList, err:= Rdb.LRange(ctx, keyname, 0, Rdb.LLen(ctx, keyname).Val()).Result()
		return jsonList, err
	}
	return nil, nil    // keyname does not exist or stores nothing
}

func ClearTaskList(ctx context.Context, userId int64, toUserId int64) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := Rdb.LTrim(ctx, keyname, 1, 0).Err()
	return err
}

func AddNewTask(ctx context.Context, userId int64, toUserId int64, taskId string) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := Rdb.LPush(ctx, keyname, taskId).Err()
	Rdb.Expire(ctx, keyname, 3600) 
	return err
}

func DeleteTask(ctx context.Context, userId int64, toUserId int64, taskId string) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := Rdb.LRem(ctx, keyname, 0, taskId).Err()
	return err
}