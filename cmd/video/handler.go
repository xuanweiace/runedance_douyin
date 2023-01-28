package main

import (
	"context"
	"encoding/hex"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	videostorage "runedance_douyin/kitex_gen/videoStorage"
	"strconv"
	"strings"
)

// VideoStorageServiceImpl implements the last service interface defined in the IDL.
type VideoStorageServiceImpl struct{}

// UploadVideoToDB implements the VideoStorageServiceImpl interface.
func getVideoDB() (string, error) {
	u := "mongodb://user02:User02@qwq.bogo.ac.cn:23317/fdb02"
	return u, nil
}
func (s *VideoStorageServiceImpl) UploadVideoToDB(ctx context.Context, req *videostorage.VideoUploadRequest) (resp *videostorage.VideoUploadResponse, err error) {

	url, err1 := getVideoDB()
	client, err2 := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err1 != nil || err2 != nil {
		ret := videostorage.VideoUploadResponse{
			Result_: false,
			Message: "mongodb unavailable",
		}
		if err1 != nil {
			err1 = err2
		}
		return &ret, err1
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			//TODO 打印日志并报告数据库关闭失败
		}
	}()

	var builder strings.Builder
	builder.WriteString(strconv.FormatInt(req.VideoId, 10))
	builder.WriteString("bogo's salty")
	hhh.Write([]byte(builder.String()))
	sid := hex.EncodeToString(hhh.Sum(nil))
	fname := strconv.FormatInt(req.VideoId, 10)

	db := client.Database("fdb02")
	err3 := UploadFileToDB(db, "fs_cover", &req.CoverData, &sid, &fname)
	if err3 != nil {
		return &videostorage.VideoUploadResponse{
			Result_: false,
			Message: "Gridfs error",
		}, err3
	}

	err4 := UploadFileToDB(db, "fs_video", &req.VideoData, &sid, &fname)
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
		return &videostorage.VideoUploadResponse{
			Result_: false,
			Message: "Gridfs error",
		}, err4
	}
	return &videostorage.VideoUploadResponse{
		Result_: true,
	}, nil
}
