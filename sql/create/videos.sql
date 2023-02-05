/*
type Video struct {
	VideoId       int64     `gorm:"primaryKey"`
	AuthorId      int64     `gorm:"index; not null"`
	CreatedAt     time.Time `gorm:"not null"`
	StorageId     string    `gorm:"size:64; not null"`
	FavoriteCount int       `gorm:"not null"`
	CommentCount  int       `gorm:"not null"`
	Title         string    `gorm:"not null"`
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
  `created_at` datetime(3) NOT NULL,
  `storage_id` varchar(64) NOT NULL,
  `favorite_count` bigint NOT NULL,
  `comment_count` bigint NOT NULL,
  `title` longtext NOT NULL,
  PRIMARY KEY (`video_id`),
  KEY `idx_videos_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
