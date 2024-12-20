DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `community_id` int(10) NOT NULL,
                        `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
                        `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_community_name` (`community_name`) USING BTREE,
                        UNIQUE KEY `idx_community_id` (`community_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



INSERT INTO `community` (`community_id`, `community_name`, `introduction`) VALUES
                                                                               (1, 'GO', 'A community for Go language enthusiasts'),
                                                                               (2, 'leetcode', 'A platform for coding practice and competitive programming'),
                                                                               (3, 'CS:GO', 'A community for Counter-Strike: Global Offensive players'),
                                                                               (4, 'LOL', 'A community for League of Legends players');


DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
    `community_id` bigint(20) NOT NULL  COMMENT '所属社区',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_name` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


UPDATE post
SET community_id = 4
WHERE community_id = 2003;
