package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	videostorage "runedance_douyin/kitex_gen/videoStorage"
	"strconv"
	"strings"
)

//	func getVideoInfo(ctx context.Context, vid int64) (*Video, error) {
//		rst, err := existInRedis(ctx, vid)
//		if err != nil {
//			return nil, err
//		}
//		if rst == 1 {
//			return queryVideoFromRedis(ctx, vid)
//		} else {
//			return queryVideoFromMysql(vid)
//		}
//	}
func getVideoInfo(ctx context.Context, vid int64) (*Video, error) {
	i, e := existInRedis(ctx, vid)
	if e == nil && i == 1 {
		return queryVideoFromRedis(ctx, vid)
	} else {
		//fmt.Println("qvfm")
		return queryVideoFromMysql(vid)
	}
}

//func pushIn(ctx context.Context, vid int64) error {
//	rst := redisClient.ZAdd(ctx, "newVideoSet", redis.Z{
//		Score:  float64(vid),
//		Member: vid,
//	})
//	return rst.Err()
//}
//func pushOut(ctx context.Context) error {
//	rst, err := redisClient.ZCount(ctx, "newVideoSet", "-inf", "+inf").Result()
//	if err != nil {
//		return err
//	}
//	if rst < constants.VideoFeedSize {
//		return nil
//	}
//
//	rst2, err2 := redisClient.ZPopMin(ctx, "newVideoSet").Result()
//	if err2 != nil {
//		return err2
//	}
//
//	v, e := queryVideoFromRedis(ctx, int64(rst2[0].Score))
//	if e != nil {
//		log.WithFields(log.Fields{
//			"request_id": ctx.Value("request_id"),
//		}).Error("推出缓存时redis查询失败")
//		return e
//	}
//
//	e1 := updateVideoInMysql(v.VideoId, "favorite_count", v.FavoriteCount)
//	e2 := updateVideoInMysql(v.VideoId, "comment_count", v.CommentCount)
//	if e1 != nil || e2 != nil {
//		log.WithFields(log.Fields{
//			"request_id": ctx.Value("request_id"),
//		}).Error("推出缓存时更新mysql失败")
//		if e1 != nil {
//			return e1
//		} else {
//			return e2
//		}
//	}
//
//	err = deleteVideoInRedis(ctx, v.VideoId)
//	if err != nil {
//		log.WithFields(log.Fields{
//			"request_id": ctx.Value("request_id"),
//		}).Error("持久化到mysql后无法从redis中删除")
//	}
//	return nil
//}

func uploadVideo(ctx context.Context, video Video, nativeData []byte) (string, error) {

	err1 := insertVideoToMysql(&video)
	if err1 != nil {
		return "Mysql Error", err1
	}

	hhh := sha256.New()

	var builder strings.Builder
	builder.WriteString(strconv.FormatInt(video.VideoId, 10))
	builder.WriteString("bogo's salty")
	hhh.Write([]byte(builder.String()))
	sid := hex.EncodeToString(hhh.Sum(nil))
	video.StorageId = sid
	_ = updateVideoInMysql(video.VideoId, "storage_id", sid)
	err1 = upToCos(video.VideoId, nativeData)
	if err1 != nil {
		e := deleteVideoInMysql(ctx, video.VideoId)
		if e != nil {
			log.WithFields(log.Fields{
				"request_id": ctx.Value("request_id"),
				"user_id":    video.AuthorId,
			}).Error("上传cos失败后无法在mysql中删除")
		}
		return "上传cos失败", err1
	}
	//time.Sleep(time.Second * 20)
	_, _ = storageClient.UploadVideoToDB(ctx, &videostorage.VideoStorageUploadRequest{
		VideoId: video.VideoId,
	})
	return "ok", nil
	/*
		if err2 != nil {
			e := deleteVideoInMysql(ctx, video.VideoId)
			if e != nil {
				log.WithFields(log.Fields{
					"request_id": ctx.Value("request_id"),
					"user_id":    video.AuthorId,
				}).Error("上传mongodb失败后无法在mysql中删除")
			}
			return "StorageService Error", err2
		}
		video.StorageId = sid
		err4 := updateVideoInMysql(video.VideoId, "storage_id", sid)
		if err4 != nil {
			log.WithFields(log.Fields{
				"request_id": ctx.Value("request_id"),
				"user_id":    video.AuthorId,
			}).Error("上传mongodb失败后无法在mysql中更新storage_id")
			return "视频文件上传成功; 存储编号转存失败", err4
		}
	*/
	//err5 := pushOut(ctx)
	//if err5 != nil {
	//	return "缓存队列满且推出失败", err5
	//}
	//err6 := insertVideoToRedis(ctx, video)
	//err7 := pushIn(ctx, video.VideoId)
	//if err6 != nil || err7 != nil {
	//	info := "上传成功; 缓存写入失败"
	//	log.WithFields(log.Fields{
	//		"request_id": ctx.Value("request_id"),
	//		"user_id":    video.AuthorId,
	//	}).Error(info)
	//	if err6 == nil {
	//		return info, err7
	//	} else {
	//		return info, err6
	//	}
	//}
	//return "ok", nil
}
func getVideoList(ctx context.Context, aid int64) (*[]int64, error) {
	if aid == 0 {
		return queryRecommendList(ctx)
		//return queryListFromRedis(ctx)
	} else {
		return queryListFromMysql(ctx, aid)
	}
}
func updateVideo[T any](ctx context.Context, vid int64, which string, what T) error {
	r, e := existInRedis(ctx, vid)
	if e == nil && r == 1 {
		return updateVideoInRedis(ctx, vid, which, what)
	} else {
		return updateVideoInMysql(vid, which, what)
	}
}
