package db_redis

import (
	"context"
	"runedance_douyin/pkg/tools"
	"time"
)

func GetPendingTaskIDs(ctx context.Context, userId int64, toUserId int64) ([]string, error){
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	// check if keyname exist in redis
	jsonList, err:= RdbCluster.LRange(ctx, keyname, 0, RdbCluster.LLen(ctx, keyname).Val()).Result()
	return jsonList, err
}

func ClearTaskList(ctx context.Context, userId int64, toUserId int64) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := RdbCluster.LTrim(ctx, keyname, 1, 0).Err()
	return err
}

func AddNewTask(ctx context.Context, userId int64, toUserId int64, taskId string) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := RdbCluster.LPush(ctx, keyname, taskId).Err()
	RdbCluster.ExpireAt(ctx, keyname, time.Now().Add(time.Hour))    // set expire time 
	return err
}

func DeleteTask(ctx context.Context, userId int64, toUserId int64, taskId string) error {
	keyname := tools.GenerateKeyname(userId, toUserId) + "-task"
	err := RdbCluster.LRem(ctx, keyname, 0, taskId).Err()
	return err
}