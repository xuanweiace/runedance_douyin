package main

import (
	"context"
	"runedance_douyin/kitex_gen/videoProcess"
	videostorage "runedance_douyin/kitex_gen/videoStorage"
	"time"
)

// VideoProcessServiceImpl implements the last service interface defined in the IDL.
type VideoProcessServiceImpl struct{}

type Video struct {
	VideoId       int64     `gorm:"primaryKey"`
	AuthorId      int64     `gorm:"index; not null"`
	CreatedAt     time.Time `gorm:"not null"`
	DeletedAt     time.Time `gorm:"not null"`
	StorageId     string    `gorm:"size:64; not null"`
	FavoriteCount int       `gorm:"not null"`
	CommentCount  int       `gorm:"not null"`
	Title         string    `gorm:"not null"`
}

// GetVideoInfo implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) GetVideoInfo(ctx context.Context, req *videoprocess.VideoInfoRequest) (resp *videoprocess.VideoInfoResponse, err error) {
	var v Video
	gormClient.First(&v, req.VideoId)
	resp = &videoprocess.VideoInfoResponse{
		VideoId:       v.VideoId,
		AuthorId:      v.AuthorId,
		StorageId:     v.StorageId,
		FavoriteCount: int32(v.FavoriteCount),
		CommentCount:  int32(v.CommentCount),
		Title:         v.Title,
	}
	//TODO 错误处理
	return resp, nil
}

// UploadVideo implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) UploadVideo(ctx context.Context, req *videoprocess.VideoProcessUploadRequest) (resp *videoprocess.VideoProcessUploadResponse, err error) {
	v := Video{
		AuthorId:      int64(req.AuthorId),
		StorageId:     "0",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	result := gormClient.Create(&v)
	if result.Error != nil {
		return &videoprocess.VideoProcessUploadResponse{
			Result_: false,
			Message: "Mysql Error",
		}, result.Error
	}

	err1 := up(v.VideoId, req.VideoData)
	if err1 != nil {
		gormClient.Delete(&v)
		return &videoprocess.VideoProcessUploadResponse{
			Result_: false,
			Message: "Cos Error",
		}, err1
	}
	time.Sleep(time.Second * 10)
	sid, err2 := storageClient.UploadVideoToDB(context.TODO(), &videostorage.VideoStorageUploadRequest{VideoId: v.VideoId})
	if err2 != nil {
		return &videoprocess.VideoProcessUploadResponse{
			Result_: false,
			Message: "Storage Server Error",
		}, err2
	} else {
		gormClient.Model(&Video{}).Where("video_id=?", v.VideoId).Update("storage_id", sid)
		return &videoprocess.VideoProcessUploadResponse{
			Result_: true,
			Message: "Accepted",
		}, nil
	}
}
func (s *VideoProcessServiceImpl) GetVideoList(ctx context.Context, authorId int64) (resp *videoprocess.VideoListResponse, err error) {
	gormClient.Model(&Video{}).Select("video_id").Where("author_id = ?", authorId).Find(&resp.Published)
	return resp, nil
}
