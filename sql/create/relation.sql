CREATE TABLE `douyin`.`relation`
(
    `fans_id` bigint NOT NULL COMMENT '关注者的用户id',
    `user_id` bigint NOT NULL COMMENT '被关注者的用户id',
    PRIMARY KEY (`fans_id`, `user_id`)
);