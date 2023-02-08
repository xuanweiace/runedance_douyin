
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
                            `Uid` bigint DEFAULT NULL,
                            `Vid` bigint DEFAULT NULL,
                            `Action` int DEFAULT NULL,
                            `Id` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                            PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;