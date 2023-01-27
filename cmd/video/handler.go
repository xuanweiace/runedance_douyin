package main

import (
	"context"
	videostorage "runedance_douyin/cmd/video/kitex_gen/videoStorage"
)

// VideoStorageServiceImpl implements the last service interface defined in the IDL.
type VideoStorageServiceImpl struct{}

// QueryVideoURL implements the VideoStorageServiceImpl interface.
func (s *VideoStorageServiceImpl) QueryVideoURL(ctx context.Context, req *videostorage.VideoURLRequest) (resp *videostorage.VideoURLResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadVideoToDB implements the VideoStorageServiceImpl interface.
func (s *VideoStorageServiceImpl) UploadVideoToDB(ctx context.Context, req *videostorage.VideoUploadRequest) (resp *videostorage.VideoUploadResponse, err error) {
	// TODO: Your code here...
	return
}
