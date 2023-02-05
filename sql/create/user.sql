-- Active: 1675424026470@@127.0.0.1@3306@douyin
CREATE TABLE `douyin`.`user`
(
    `user_id` bigint AUTO_INCREMENT,
    `username` VARCHAR(18) UNIQUE,
    `password` VARCHAR(18) NOT NULL,
    `avatar` VARCHAR(18) ,
    `Salt` VARCHAR(18),
    `follow_count` bigint,
    `follower_count` bigint,
    PRIMARY KEY (user_id)
);