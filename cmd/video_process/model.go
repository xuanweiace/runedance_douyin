package main

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	VideoId       int64 `gorm:"primaryKey" redis:"video_id"`
	AuthorId      int64 `gorm:"index; not null" redis:"author_id"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	StorageId     string `gorm:"size:64; not null" redis:"storage_id"`
	FavoriteCount int32  `gorm:"not null" redis:"favorite_count"`
	CommentCount  int32  `gorm:"not null" redis:"comment_count"`
	Title         string `gorm:"not null" redis:"title"`
}
