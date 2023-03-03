package main

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"runedance_douyin/kitex_gen/videoStorage"
	"strconv"
	"time"
)

// VideoStorageServiceImpl implements the last service interface defined in the IDL.
type VideoStorageServiceImpl struct{}

// UploadVideoToDB implements the VideoStorageServiceImpl interface.
func (s *VideoStorageServiceImpl) UploadVideoToDB(ctx context.Context, req *videostorage.VideoStorageUploadRequest) (string, error) {

	time.Sleep(time.Second * 5)

	sid := generateSID(req.VideoId)
	fname := strconv.FormatInt(req.VideoId, 10)
	var videoData, coverData *[]byte
	var err error
	for t := 3; t != 0; t-- {
		time.Sleep(time.Second * 5)
		videoData, coverData, err = DownloadFromCos(fname)
		if err == nil {
			break
		}
	}
	if err != nil {

	}
	db := mongoClient.Database("fdb02")
	err3 := UploadFileToDB(db, "fs_cover", coverData, &sid, &fname)
	if err3 != nil {
		return "", err3
	}
	err4 := UploadFileToDB(db, "fs_video", videoData, &sid, &fname)
	if err4 != nil { //视频文件未成功上传则删除封面图
		id, e1 := primitive.ObjectIDFromHex("sid")
		opts := options.GridFSBucket().SetName("fs_cover")
		bucket, e2 := gridfs.NewBucket(db, opts)
		e3 := errors.New("flag")
		if e2 != nil && e1 != nil {
			e3 = bucket.Delete(id)
		}
		if e3 != nil { //未成功删除封面图
			log.WithFields(log.Fields{
				"request_id": ctx.Value("request_id"),
			}).Error("MongoDB Error")
		}
		return "", err4
	}
	err5 := updateVideoInMysql(req.VideoId, "storage_id", sid)
	return sid, err5
}
