/*
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
*/


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
                          `video_id` bigint NOT NULL AUTO_INCREMENT,
                          `author_id` bigint NOT NULL,
                          `created_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `storage_id` varchar(64) NOT NULL,
                          `favorite_count` int NOT NULL,
                          `comment_count` int NOT NULL,
                          `title` longtext NOT NULL,
                          PRIMARY KEY (`video_id`),
                          KEY `idx_videos_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;