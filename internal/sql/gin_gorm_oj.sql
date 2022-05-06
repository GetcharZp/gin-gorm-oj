/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.8
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 192.168.1.8:3306
 Source Schema         : gin_gorm_oj

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 05/05/2022 20:58:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category_basic
-- ----------------------------
DROP TABLE IF EXISTS `category_basic`;
CREATE TABLE `category_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL COMMENT '分类名称',
  `parent_id` int(11) DEFAULT '0' COMMENT '父级ID',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for problem_basic
-- ----------------------------
DROP TABLE IF EXISTS `problem_basic`;
CREATE TABLE `problem_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL COMMENT '问题的题目',
  `content` text COMMENT '问题的正文描述',
  `max_runtime` int(11) DEFAULT NULL COMMENT '最大的运行时间',
  `max_mem` int(11) DEFAULT NULL COMMENT '最大的运行内存',
  `pass_num` int(11) DEFAULT '0' COMMENT '通过的问题个数',
  `submit_num` int(11) DEFAULT '0' COMMENT '提交次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for problem_category
-- ----------------------------
DROP TABLE IF EXISTS `problem_category`;
CREATE TABLE `problem_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `problem_id` int(11) DEFAULT NULL COMMENT '问题ID',
  `category_id` int(11) DEFAULT NULL COMMENT '分类ID',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for submit_basic
-- ----------------------------
DROP TABLE IF EXISTS `submit_basic`;
CREATE TABLE `submit_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `problem_identity` varchar(36) DEFAULT NULL COMMENT '问题的唯一标识',
  `user_identity` varchar(36) DEFAULT NULL COMMENT '用户的唯一标识',
  `path` varchar(255) DEFAULT NULL COMMENT '代码路径',
  `status` tinyint(1) DEFAULT '-1' COMMENT '【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存】',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for test_case
-- ----------------------------
DROP TABLE IF EXISTS `test_case`;
CREATE TABLE `test_case` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) NOT NULL,
  `problem_identity` varchar(36) DEFAULT NULL,
  `input` text,
  `output` text,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL COMMENT '唯一标识',
  `name` varchar(100) DEFAULT NULL COMMENT '名称',
  `password` varchar(32) DEFAULT NULL COMMENT '密码',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `mail` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `pass_num` int(11) DEFAULT '0' COMMENT '完成问题的个数',
  `submit_num` int(11) DEFAULT '0' COMMENT '总提交次数',
  `is_admin` tinyint(1) DEFAULT '0' COMMENT '是否是管理员【0-否，1-是】',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
