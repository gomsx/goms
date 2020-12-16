/*
 Navicat Premium Data Transfer

 Source Server         : gomstest
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : 192.168.43.204:3306
 Source Schema         : test_db

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 10/03/2020 01:15:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ping_table
-- ----------------------------
DROP TABLE IF EXISTS `ping_table`;
CREATE TABLE `ping_table`  (
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `count` int(20) UNSIGNED ZEROFILL NOT NULL,
  PRIMARY KEY (`type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ping_table
-- ----------------------------
INSERT INTO `ping_table` VALUES ('grpc', 00000000000000000001);
INSERT INTO `ping_table` VALUES ('http', 00000000000000000002);

-- ----------------------------
-- Table structure for user_table
-- ----------------------------
DROP TABLE IF EXISTS `user_table`;
CREATE TABLE `user_table`  (
  `uid` bigint(20) NOT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `sex` int(10) NULL DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_table
-- ----------------------------
INSERT INTO `user_table` VALUES (3954, 'xxx', 1);
INSERT INTO `user_table` VALUES (205140319971954, 'yyy', 1);

SET FOREIGN_KEY_CHECKS = 1;

