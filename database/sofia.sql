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

 Date: 14/12/2020 21:09:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for answers
-- ----------------------------
DROP TABLE IF EXISTS `answers`;
CREATE TABLE `answers`  (
  `aid` bigint(20) NOT NULL AUTO_INCREMENT,
  `answerer` bigint(20) NOT NULL,
  `qid` bigint(20) NOT NULL,
  `comment_count` bigint(20) NOT NULL,
  `criticism_count` bigint(20) NOT NULL,
  `like_count` bigint(20) NOT NULL,
  `approval_count` bigint(20) NOT NULL,
  `time` bigint(20) NOT NULL,
  PRIMARY KEY (`aid`) USING BTREE,
  INDEX `answerer`(`answerer`) USING BTREE,
  INDEX `qid`(`qid`) USING BTREE,
  CONSTRAINT `answers_ibfk_1` FOREIGN KEY (`answerer`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `answers_ibfk_2` FOREIGN KEY (`qid`) REFERENCES `questions` (`qid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for approve_answers
-- ----------------------------
DROP TABLE IF EXISTS `approve_answers`;
CREATE TABLE `approve_answers`  (
  `uid` bigint(20) NOT NULL,
  `aid` bigint(20) NOT NULL,
  PRIMARY KEY (`uid`, `aid`) USING BTREE,
  INDEX `aid`(`aid`) USING BTREE,
  CONSTRAINT `approve_answers_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `approve_answers_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `answers` (`aid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `cmid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL,
  `aid` bigint(20) NOT NULL,
  `time` bigint(20) NOT NULL,
  PRIMARY KEY (`cmid`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `aid`(`aid`) USING BTREE,
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `answers` (`aid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for criticisms
-- ----------------------------
DROP TABLE IF EXISTS `criticisms`;
CREATE TABLE `criticisms`  (
  `ctid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL,
  `aid` bigint(20) NOT NULL,
  `time` bigint(20) NOT NULL,
  PRIMARY KEY (`ctid`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `aid`(`aid`) USING BTREE,
  CONSTRAINT `criticisms_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `criticisms_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `answers` (`aid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorite_items
-- ----------------------------
DROP TABLE IF EXISTS `favorite_items`;
CREATE TABLE `favorite_items`  (
  `fid` bigint(20) NOT NULL,
  `qid` bigint(20) NOT NULL,
  PRIMARY KEY (`fid`, `qid`) USING BTREE,
  INDEX `favorite_items_ibfk_2`(`qid`) USING BTREE,
  CONSTRAINT `favorite_items_ibfk_1` FOREIGN KEY (`fid`) REFERENCES `favorites` (`fid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `favorite_items_ibfk_2` FOREIGN KEY (`qid`) REFERENCES `questions` (`qid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `fid` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL,
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`fid`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  CONSTRAINT `favorites_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for followers
-- ----------------------------
DROP TABLE IF EXISTS `followers`;
CREATE TABLE `followers`  (
  `uid` bigint(20) NOT NULL,
  `follower` bigint(20) NOT NULL,
  PRIMARY KEY (`uid`, `follower`) USING BTREE,
  INDEX `follower`(`follower`) USING BTREE,
  CONSTRAINT `followers_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `followers_ibfk_2` FOREIGN KEY (`follower`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hotlist_items
-- ----------------------------
DROP TABLE IF EXISTS `hotlist_items`;
CREATE TABLE `hotlist_items`  (
  `qid` bigint(20) NOT NULL,
  `temperature` bigint(20) NOT NULL,
  PRIMARY KEY (`qid`) USING BTREE,
  CONSTRAINT `hotlist_items_ibfk_1` FOREIGN KEY (`qid`) REFERENCES `questions` (`qid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for labels
-- ----------------------------
DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels`  (
  `lid` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`lid`) USING BTREE,
  UNIQUE INDEX `title`(`title`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for like_answers
-- ----------------------------
DROP TABLE IF EXISTS `like_answers`;
CREATE TABLE `like_answers`  (
  `uid` bigint(20) NOT NULL,
  `aid` bigint(20) NOT NULL,
  PRIMARY KEY (`uid`, `aid`) USING BTREE,
  INDEX `aid`(`aid`) USING BTREE,
  CONSTRAINT `like_answers_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `like_answers_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `answers` (`aid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for question_labels
-- ----------------------------
DROP TABLE IF EXISTS `question_labels`;
CREATE TABLE `question_labels`  (
  `qid` bigint(20) NOT NULL,
  `lid` bigint(20) NOT NULL,
  PRIMARY KEY (`qid`, `lid`) USING BTREE,
  INDEX `question_labels_ibfk_2`(`lid`) USING BTREE,
  CONSTRAINT `question_labels_ibfk_1` FOREIGN KEY (`qid`) REFERENCES `questions` (`qid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `question_labels_ibfk_2` FOREIGN KEY (`lid`) REFERENCES `labels` (`lid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for questions
-- ----------------------------
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions`  (
  `qid` bigint(20) NOT NULL AUTO_INCREMENT,
  `raiser` bigint(20) NOT NULL,
  `category` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `accepted_answer` bigint(20) NULL DEFAULT NULL,
  `answer_count` bigint(20) NOT NULL,
  `view_count` bigint(20) NOT NULL,
  `favorite_count` bigint(20) NOT NULL,
  `time` bigint(20) NOT NULL,
  `scanned` tinyint(4) NOT NULL,
  PRIMARY KEY (`qid`) USING BTREE,
  INDEX `questions_ibfk_1`(`raiser`) USING BTREE,
  INDEX `questions_ibfk_2`(`accepted_answer`) USING BTREE,
  CONSTRAINT `questions_ibfk_1` FOREIGN KEY (`raiser`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `questions_ibfk_2` FOREIGN KEY (`accepted_answer`) REFERENCES `answers` (`aid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_labels
-- ----------------------------
DROP TABLE IF EXISTS `user_labels`;
CREATE TABLE `user_labels`  (
  `uid` bigint(20) NOT NULL,
  `lid` bigint(20) NOT NULL,
  PRIMARY KEY (`uid`, `lid`) USING BTREE,
  INDEX `user_labels_ibfk_2`(`lid`) USING BTREE,
  CONSTRAINT `user_labels_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`uid`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `user_labels_ibfk_2` FOREIGN KEY (`lid`) REFERENCES `labels` (`lid`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT,
  `oid` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `salt` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `hash_password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `gender` tinyint(4) NOT NULL,
  `role` tinyint(4) NOT NULL,
  `account_type` tinyint(4) NOT NULL,
  `active_code` bigint(20) NOT NULL,
  `passwd_code` bigint(20) NOT NULL,
  `exp` bigint(20) NOT NULL,
  `follower_count` bigint(20) NOT NULL,
  `following_count` bigint(20) NOT NULL,
  `notification_time` bigint(20) NOT NULL,
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
