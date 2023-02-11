package main

import (
	"context"
	"encoding/hex"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"runedance_douyin/kitex_gen/videoStorage"
	"strconv"
	"strings"
)

// VideoStorageServiceImpl implements the last service interface defined in the IDL.
type VideoStorageServiceImpl struct{}

// UploadVideoToDB implements the VideoStorageServiceImpl interface.
func (s *VideoStorageServiceImpl) UploadVideoToDB(ctx context.Context, req *videostorage.VideoStorageUploadRequest) (string, error) {

	var builder strings.Builder
	builder.WriteString(strconv.FormatInt(req.VideoId, 10))
	builder.WriteString("bogo's salty")
	hhh.Write([]byte(builder.String()))
	sid := hex.EncodeToString(hhh.Sum(nil))
	fname := strconv.FormatInt(req.VideoId, 10)
	db := mongoClient.Database("fdb02")
	videoData, coverData := DownloadFromCos(fname)
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
			//TODO 打印日志并报告rollback失败
		}
		return "", err4
	}
	return sid, nil
}
