package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	videostorage "runedance_douyin/kitex_gen/videoStorage"
)

func getVideoInfo(ctx context.Context, vid int64) (*Video, error) {
	i, e := existInRedis(ctx, vid)
	if e == nil && i == 1 {
		return queryVideoFromRedis(ctx, vid)
	} else {
		//fmt.Println("qvfm")
		v, e := queryVideoFromMysql(ctx, vid)
		err := insertVideoToRedis(ctx, *v)
		if err != nil {
			log.WithFields(log.Fields{
				"request_id": ctx.Value("request_id"),
			}).Error("无法缓存视频信息")
		}
		return v, e
	}
}

func uploadVideo(ctx context.Context, video Video, nativeData []byte) (string, error) {

	err1 := insertVideoToMysql(ctx, &video)
	if err1 != nil {
		return "Mysql Error", err1
	}

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
	//log.Println("update video")
	//log.Println(vid, which, what)
	//log.Println(r)
	//log.Println(e)
	if e == nil && r == 1 {
		return updateVideoInRedis(ctx, vid, which, what)
	} else {
		return updateVideoInMysql(ctx, vid, which, what)
	}
}
