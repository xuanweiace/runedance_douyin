CREATE TABLE `douyin`.`user`
(
    `user_id` bigint AUTO_INCREMENT,
    `username` VARCHAR(18) UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `avatar` VARCHAR(18) ,
    `Salt` VARCHAR(18),
    `follow_count` bigint,
    `follower_count` bigint,
    PRIMARY KEY (user_id)
);