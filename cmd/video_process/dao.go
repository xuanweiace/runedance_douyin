package main

import (
	"bytes"
	"context"
	"github.com/gomodule/redigo/redis"
	constants "runedance_douyin/pkg/consts"
	"strconv"
)

func upToCos(id int64, file []byte) error {
	fname := "/dir01/" + strconv.FormatInt(id, 10) + ".mp4"
	_, err := cosClient.Object.Put(context.Background(), fname, bytes.NewReader(file), nil)
	return err
}

func existInRedis(ctx context.Context, vid int64) (int64, error) {
	rst := redisClient.Exists(ctx, strconv.FormatInt(vid, 10))
	return rst.Result()
}
func queryVideoFromRedis(ctx context.Context, vid int64) (*Video, error) {
	var v Video
	tmp := redisClient.HGetAll(ctx, strconv.FormatInt(vid, 10))
	values, err1 := redis.Values(tmp.Result())
	if err1 != nil {
		return nil, err1
	}
	err2 := redis.ScanStruct(values, &v)
	if err2 != nil {
		return nil, err2
	} else {
		return &v, nil
	}
}

func queryListFromRedis(ctx context.Context) (*[]int64, error) {
	rst := redisClient.Keys(ctx, "*")
	if rst.Err() != nil {
		return nil, rst.Err()
	}
	var vids []int64
	for _, v := range rst.Val() {
		i, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			return nil, e
		}
		vids = append(vids, i)
	}
	return &vids, nil
}

//	func pop(ctx context.Context) (int64, error) {
//		rst, err := redisClient.ZCount(ctx, "newVideoSet", "-inf", "+inf").Result()
//		if err != nil {
//			return -1, err
//		}
//		if rst < constants.VideoFeedSize {
//			return 0, nil
//		}
//		rst2, err2 := redisClient.ZPopMin(ctx, "newVideoSet").Result()
//		if err2 != nil {
//			return -1, err2
//		}
//		return int64(rst2[0].Score), nil
//	}
func queryVideoFromMysql(vid int64) (*Video, error) {
	var v Video
	result := gormClient.First(&v, vid)
	if result.Error != nil {
		return &v, nil
	} else {
		return nil, result.Error
	}
}
func insertVideoToRedis(ctx context.Context, video Video) error {
	//redisClient.Do(context.Background(), "hmset", redis.Args{"Video"}.AddFlat(video))
	//rst := redisClient.Do(ctx, redis.Args{video.VideoId}.AddFlat(video)...)
	//return rst.Err()
	_, err := redisClient.HMSet(ctx, strconv.FormatInt(video.VideoId, 10), video).Result()
	return err
}
func insertVideoToMysql(video *Video) error {
	result := gormClient.Create(video)
	return result.Error
}
func updateVideoInMysql[T any](vid int64, columnName string, newValue T) error {
	result := gormClient.Model(&Video{}).Where("video_id=?", vid).Update(columnName, newValue)
	return result.Error
}
func updateVideoInRedis[T any](ctx context.Context, vid int64, propertyName string, newValue T) error {
	rst := redisClient.HSet(ctx, strconv.FormatInt(vid, 10), propertyName, newValue)
	return rst.Err()
}
func deleteVideoInRedis(ctx context.Context, vid int64) error {
	rst := redisClient.Del(ctx, strconv.FormatInt(vid, 10))
	return rst.Err()
}
func deleteVideoInMysql(ctx context.Context, vid int64) error {
	rst := gormClient.Delete(&Video{}, vid)
	return rst.Error
}
func queryListFromMysql(ctx context.Context, aid int64) (*[]int64, error) {
	var published []int64
	rst := gormClient.Model(&Video{}).Limit(constants.VideoFeedSize).Where("author_id = ?", aid).Select("video_id").Find(&published)
	if rst.Error != nil {
		return nil, rst.Error
	} else {
		return &published, nil
	}
}
