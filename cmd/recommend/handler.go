package main

import (
	"context"
	recommend "runedance_douyin/cmd/recommend/kitex_gen/recommend"
	"time"
)

// RecommendServiceImpl implements the last service interface defined in the IDL.
type RecommendServiceImpl struct{}

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

// GetRecommended implements the RecommendServiceImpl interface.
func (s *RecommendServiceImpl) GetRecommended(ctx context.Context, user int64) (resp *recommend.RecommendResponse, err error) {
	gormClient.Model(&Video{}).Order("created_at").Select("video_id").Limit(30).Find(&resp.Recommended)
	return resp, nil
}
