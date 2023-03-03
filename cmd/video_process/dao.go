package main

import (
	"bytes"
	"context"
	constants "runedance_douyin/pkg/consts"
	"strconv"
	"time"
)

func upToCos(id int64, file []byte) error {
	fname := "/dir01/" + strconv.FormatInt(id, 10) + ".mp4"
	_, err := cosClient.Object.Put(context.Background(), fname, bytes.NewReader(file), nil)
	return err
}

func existInRedis(ctx context.Context, vid int64) (int64, error) {
	rst := redisClient.Exists(ctx, strconv.FormatInt(vid, 10))
	return rst.Val(), rst.Err()
}
func queryVideoFromRedis(ctx context.Context, vid int64) (*Video, error) {
	var v Video
	res := redisClient.HGetAll(context.TODO(), strconv.FormatInt(vid, 10))
	if res.Err() != nil {
		return nil, res.Err()
	}
	err1 := res.Scan(&v)
	if err1 != nil {
		return nil, err1
	} else {
		return &v, nil
	}
}

//func queryListFromRedis(ctx context.Context) (*[]int64, error) {
//	rst := redisClient.Keys(ctx, "*")
//	if rst.Err() != nil {
//		return nil, rst.Err()
//	}
//	var vids []int64
//	for _, v := range rst.Val() {
//		i, e := strconv.ParseInt(v, 10, 64)
//		if e != nil {
//			return nil, e
//		}
//		vids = append(vids, i)
//	}
//	return &vids, nil
//}

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
func queryVideoFromMysql(ctx context.Context, vid int64) (*Video, error) {
	var v Video
	v.VideoId = vid
	result := gormClient.WithContext(ctx).First(&v)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &v, nil
	}
}
func insertVideoToRedis(ctx context.Context, video Video) error {
	key := strconv.FormatInt(video.VideoId, 10)
	_, err := redisClient.HMSet(ctx, key, video).Result()
	redisClient.PExpire(ctx, key, time.Minute*10)
	return err
}
func insertVideoToMysql(ctx context.Context, video *Video) error {
	result := gormClient.WithContext(ctx).Create(video)
	return result.Error
}
func updateVideoInMysql[T any](ctx context.Context, vid int64, columnName string, newValue T) error {
	result := gormClient.WithContext(ctx).Model(&Video{}).Where("video_id=?", vid).Update(columnName, newValue)
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
	rst := gormClient.WithContext(ctx).Delete(&Video{}, vid)
	return rst.Error
}
func queryListFromMysql(ctx context.Context, aid int64) (*[]int64, error) {
	var published []int64
	rst := gormClient.WithContext(ctx).Model(&Video{}).Limit(constants.VideoFeedSize).Where("author_id = ? AND storage_id <> ?", aid, "0").Select("video_id").Find(&published)
	if rst.Error != nil {
		return nil, rst.Error
	} else {
		return &published, nil
	}
}

var cacheList []int64
var lastQuery time.Time

func queryRecommendList(ctx context.Context) (*[]int64, error) {

	current := time.Now()
	if cacheList == nil || current.Sub(lastQuery).Minutes() >= 1 {
		var published []int64
		rst := gormClient.WithContext(ctx).Model(&Video{}).Limit(constants.VideoFeedSize).Where("storage_id <> ?", "0").Order("created_at desc").Select("video_id").Find(&published)
		if rst.Error != nil {
			return nil, rst.Error
		} else {
			lastQuery = current
			cacheList = published
			return &published, nil
		}
	} else {
		return &cacheList, nil
	}
}
