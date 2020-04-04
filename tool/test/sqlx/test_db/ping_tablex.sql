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

 Date: 02/03/2020 15:21:16
*/
use dbtest;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ping_table
-- ----------------------------
DROP TABLE IF EXISTS `ping_table`;
CREATE TABLE `ping_table`  (
  `type` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `count` int(20) UNSIGNED ZEROFILL NOT NULL,
  PRIMARY KEY (`type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ping_table
-- ----------------------------
INSERT INTO `ping_table` VALUES ('grpc', 00000000000000000001);
INSERT INTO `ping_table` VALUES ('http', 00000000000000000000);

SET FOREIGN_KEY_CHECKS = 1;
