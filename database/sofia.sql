/*
 Navicat Premium Data Transfer

 Source Server         : localhost_mysql
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : sofia

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 17/11/2020 15:39:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for approve_answers
-- ----------------------------
DROP TABLE IF EXISTS `approve_answers`;
CREATE TABLE `approve_answers`  (
  `uid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `aid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`uid`, `aid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorite_items
-- ----------------------------
DROP TABLE IF EXISTS `favorite_items`;
CREATE TABLE `favorite_items`  (
  `fid` bigint(20) NOT NULL,
  `qid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`fid`, `qid`) USING BTREE,
  CONSTRAINT `favorite_items_ibfk_1` FOREIGN KEY (`fid`) REFERENCES `favorites` (`fid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `fid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`fid`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hotlist_items
-- ----------------------------
DROP TABLE IF EXISTS `hotlist_items`;
CREATE TABLE `hotlist_items`  (
  `qid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `temperature` bigint(20) NOT NULL,
  PRIMARY KEY (`qid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for kcard_attrs
-- ----------------------------
DROP TABLE IF EXISTS `kcard_attrs`;
CREATE TABLE `kcard_attrs`  (
  `kid` bigint(20) NOT NULL,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `origin` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`kid`, `name`, `value`, `origin`) USING BTREE,
  CONSTRAINT `kcard_attrs_ibfk_1` FOREIGN KEY (`kid`) REFERENCES `kcards` (`kid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for kcards
-- ----------------------------
DROP TABLE IF EXISTS `kcards`;
CREATE TABLE `kcards`  (
  `kid` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`kid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for labels
-- ----------------------------
DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels`  (
  `lid` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`lid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for like_answers
-- ----------------------------
DROP TABLE IF EXISTS `like_answers`;
CREATE TABLE `like_answers`  (
  `uid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `aid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`uid`, `aid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for question_labels
-- ----------------------------
DROP TABLE IF EXISTS `question_labels`;
CREATE TABLE `question_labels`  (
  `qid` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `lid` bigint(20) NOT NULL,
  PRIMARY KEY (`qid`, `lid`) USING BTREE,
  INDEX `lid`(`lid`) USING BTREE,
  CONSTRAINT `question_labels_ibfk_1` FOREIGN KEY (`lid`) REFERENCES `labels` (`lid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
