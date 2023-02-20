package main

import (
	"context"
	videoprocess "runedance_douyin/kitex_gen/videoProcess"
)

// VideoProcessServiceImpl implements the last service interface defined in the IDL.
type VideoProcessServiceImpl struct{}

// GetVideoInfo implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) GetVideoInfo(ctx context.Context, req *videoprocess.VideoInfoRequest) (resp *videoprocess.VideoInfoResponse, err error) {
	video, err1 := getVideoInfo(ctx, req.VideoId)
	return &videoprocess.VideoInfoResponse{
		VideoId:       video.VideoId,
		AuthorId:      video.AuthorId,
		StorageId:     video.StorageId,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	}, err1
}

// UploadVideo implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) UploadVideo(ctx context.Context, req *videoprocess.VideoProcessUploadRequest) (resp *videoprocess.VideoProcessUploadResponse, err error) {
	v := Video{
		AuthorId:      req.AuthorId,
		StorageId:     "0",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	msg, err1 := uploadVideo(ctx, v, req.VideoData)
	if err1 != nil {
		return &videoprocess.VideoProcessUploadResponse{
			Result_: false,
			Message: msg,
		}, err1
	} else {
		return &videoprocess.VideoProcessUploadResponse{
			Result_: true,
			Message: "ok",
		}, nil
	}
}

// GetVideoList implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) GetVideoList(ctx context.Context, authorId int64) (resp *videoprocess.VideoListResponse, err error) {
	r, e := getVideoList(ctx, authorId)
	if e != nil {
		return nil, e
	}
	return &videoprocess.VideoListResponse{Published: *r}, nil

}

// ChangeFavCount implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) ChangeFavCount(ctx context.Context, req *videoprocess.ChangeCountRequest) (err error) {
	return updateVideo(ctx, req.VideoId, "favorite_count", req.ChangeValue)
}

// ChangeComCount implements the VideoProcessServiceImpl interface.
func (s *VideoProcessServiceImpl) ChangeComCount(ctx context.Context, req *videoprocess.ChangeCountRequest) (err error) {
	return updateVideo(ctx, req.VideoId, "comment_count", req.ChangeValue)
}
