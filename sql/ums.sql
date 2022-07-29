/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version :
 Source Host           : localhost:3306
 Source Schema         :

 Target Server Type    : MySQL
 Target Server Version :
 File Encoding         :

 Date: 24/07/2022 00:32:28
*/

-- 这是artmall的sql文件
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -- ----------------------------
-- -- user
-- id
-- 昵称,手机号(唯一登录名)
-- 密码, 邮箱,头像, 性别,
-- 生日, 所在城市, 职业,
-- 个性签名, 是否禁用 是否删除
-- 创建/更新/删除时间
-- -- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `mobile` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `header_image` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `gender` tinyint(4) NULL DEFAULT NULL COMMENT '性别[0 - 男性，1 - 女性]',
  `birthday` date NULL DEFAULT '2005-06-08' COMMENT '生日',
  `city` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所在城市',
  `job` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职业',
  `sign` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个性签名',
  `is_disable` int(1) NULL DEFAULT 0 COMMENT '是否禁用[0-不禁用，1-禁用]',
  `is_delete` int(1) NULL DEFAULT 0 COMMENT '是否删除[0-不删除，1-删除]',
  `create_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;


-- 这是bluebell的sql文件，实际上我是直接用gorm来建表的，没用到本sql文件
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `gorm_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'gorm_id',
    `gorm_created_at` datetime(3) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `gorm_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `gorm_deleted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别[0 - 男性，1 - 女性]',

PRIMARY KEY (`gorm_id`),
UNIQUE KEY `idx_username` (`username`) USING BTREE,
UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;