SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `Id` bigint NOT NULL,
                           `Uid` bigint NOT NULL,
                           `Vid` bigint NOT NULL,
                           `Content` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `Content_date` time DEFAULT NULL,
                           PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

