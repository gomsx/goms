/*
 Navicat Premium Data Transfer

 Source Server         : gomstest
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : 192.168.43.204:3306
 Source Schema         : dbtest

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 01/03/2020 18:35:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_call_ping_count
-- ----------------------------
DROP TABLE IF EXISTS `api_call_ping_count`;
CREATE TABLE `api_call_ping_count`  (
  `type` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `count` int(20) UNSIGNED ZEROFILL NOT NULL,
  PRIMARY KEY (`type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of api_call_ping_count
-- ----------------------------
INSERT INTO `api_call_ping_count` VALUES ('grpc', 00000000000000000004);
INSERT INTO `api_call_ping_count` VALUES ('http', 00000000000000000004);

SET FOREIGN_KEY_CHECKS = 1;
