package main

import "context"

func updateFavoriteToRedis0[T any](ctx context.Context, action int32, key string) error {
	res := redisClient.Set(ctx, key, action, 1000000*1000*60*30)
	return res.Err()
}

func updateFavoriteToRedis1[T any](ctx context.Context, action int32, uid string, vid string) error {
	//1 看redis里有没有重复的
	//1 是否有uid
	//3 put进去
	res := redisClient.Exists(ctx, uid)
	_, e := res.Result()
	if e != nil {
		//设置过期时间
		redisClient.HSet(ctx, uid, vid, action)
		redisClient.Expire(ctx, uid, 1000000*1000*60)
	} else {
		redisClient.HSet(ctx, uid, vid, action)
	}
	res = redisClient.Exists(ctx, vid)
	_, err := res.Result()
	if err != nil {
		//设置过期时间
		redisClient.HSet(ctx, vid, uid, action)
		redisClient.Expire(ctx, vid, 1000000*1000*60)
	} else {
		redisClient.HSet(ctx, vid, uid, action)
	}
	return nil
}
