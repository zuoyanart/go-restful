/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50173
Source Host           : 192.168.1.117:3306
Source Database       : pizzaCms

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2016-03-04 10:53:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pz_article
-- ----------------------------
DROP TABLE IF EXISTS `pz_article`;
CREATE TABLE `pz_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT '',
  `timg` varchar(100) DEFAULT '',
  `content` varchar(10000) DEFAULT '',
  `brief` varchar(255) DEFAULT '',
  `nodeid` int(11) DEFAULT '0',
  `count` int(11) DEFAULT '0',
  `reco` int(11) DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '0',
  `pass` int(11) DEFAULT '0',
  `source` varchar(100) DEFAULT '',
  `tags` varchar(100) DEFAULT '',
  `link` varchar(100) DEFAULT '',
  `comment` int(11) DEFAULT '0',
  `state` int(11) DEFAULT '0',
  `createtime` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `page` (`id`,`title`,`nodeid`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_article
-- ----------------------------
INSERT INTO `pz_article` VALUES ('1', '战略上坚持持久战', '', '战略是谋划全局、决定长远的策略。全面建成小康社会、加快推进社会主义现代化、实现中华民族伟大复兴，都需要从战略高度审时度势地提出了治国理政的新理念新思想新举措，并进行系统的战略谋划和战略布局。2012年9月1日习近平在中央党校2012年秋季学期开学典礼上指出：“我们要全面建成小康社会和实现社会主义现代化，有许许多多重大问题需要进行战略谋划。凡是涉及我国经济、政治、文化、社会、生态、外交、国防和党的建设的全局性的重大问题，都需要从战略上进行思考、研究和筹谋；凡是涉及改革发展稳定工作中的各种重大问题，也都需要从战略上拿出治本之策。”', '面对世界经济复苏乏力、国内经济下行压力增大的复杂局面，作为世界经济发展领头羊的中国经济该拿出什么样重大举措才能再创奇迹？2015年11月10日，习近平在中央财经领导小组第十一次会议上给出答案：“战略上坚持持久战，战术上打好歼灭战。” 请随“学习中国”小编一起学习。', '1', '0', '0', '1', '0', '', '', '', '0', '0', '1');
INSERT INTO `pz_article` VALUES ('2', '习近平振兴中国经济的战略和战术1', '', '1战略是谋划全局、决定长远的策略。全面建成小康社会、加快推进社会主义现代化、实现中华民族伟大复兴，都需要从战略高度审时度势地提出了治国理政的新理念新思想新举措，并进行系统的战略谋划和战略布局。2012年9月1日习近平在中央党校2012年秋季学期开学典礼上指出：“我们要全面建成小康社会和实现社会主义现代化，有许许多多重大问题需要进行战略谋划。凡是涉及我国经济、政治、文化、社会、生态、外交、国防和党的建设的全局性的重大问题，都需要从战略上进行思考、研究和筹谋；凡是涉及改革发展稳定工作中的各种重大问题，也都需要从战略上拿出治本之策。”', '1面对世界经济复苏乏力、国内经济下行压力增大的复杂局面，作为世界经济发展领头羊的中国经济该拿出什么样重大举措才能再创奇迹？2015年11月10日，习近平在中央财经领导小组第十一次会议上给出答案：“战略上坚持持久战，战术上打好歼灭战。” 请随“学习中国”小编一起学习。', '1', '0', '0', '1', '0', '', '', '', '0', '0', '1');
INSERT INTO `pz_article` VALUES ('3', '习近平振兴中国经济的战略和战术2', '', '2战略是谋划全局、决定长远的策略。全面建成小康社会、加快推进社会主义现代化、实现中华民族伟大复兴，都需要从战略高度审时度势地提出了治国理政的新理念新思想新举措，并进行系统的战略谋划和战略布局。2012年9月1日习近平在中央党校2012年秋季学期开学典礼上指出：“我们要全面建成小康社会和实现社会主义现代化，有许许多多重大问题需要进行战略谋划。凡是涉及我国经济、政治、文化、社会、生态、外交、国防和党的建设的全局性的重大问题，都需要从战略上进行思考、研究和筹谋；凡是涉及改革发展稳定工作中的各种重大问题，也都需要从战略上拿出治本之策。”', '2面对世界经济复苏乏力、国内经济下行压力增大的复杂局面，作为世界经济发展领头羊的中国经济该拿出什么样重大举措才能再创奇迹？2015年11月10日，习近平在中央财经领导小组第十一次会议上给出答案：“战略上坚持持久战，战术上打好歼灭战。” 请随“学习中国”小编一起学习。', '1', '0', '0', '1', '0', '', '', '', '0', '0', '1');

-- ----------------------------
-- Table structure for pz_comment
-- ----------------------------
DROP TABLE IF EXISTS `pz_comment`;
CREATE TABLE `pz_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `articleid` int(11) DEFAULT '0' COMMENT '文章id',
  `addtime` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  `content` varchar(1000) DEFAULT '' COMMENT '评论内容',
  `uid` int(11) DEFAULT '0' COMMENT '用户id',
  `username` varchar(30) DEFAULT '' COMMENT '用户昵称',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_comment
-- ----------------------------

-- ----------------------------
-- Table structure for pz_node
-- ----------------------------
DROP TABLE IF EXISTS `pz_node`;
CREATE TABLE `pz_node` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0',
  `name` varchar(50) DEFAULT '',
  `brief` varchar(255) DEFAULT '',
  `nodepath` varchar(255) DEFAULT '',
  `link` varchar(100) DEFAULT '',
  `weight` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_node
-- ----------------------------
INSERT INTO `pz_node` VALUES ('1', '0', '首页', '', ',1,', '', '0');
INSERT INTO `pz_node` VALUES ('3', '1', '新闻', '', ',1,3,', '111方芳芳', '0');
INSERT INTO `pz_node` VALUES ('4', '1', '资讯', '', ',1,4,', '', '0');
INSERT INTO `pz_node` VALUES ('5', '1', 'cesces', '', ',1,5,', '', '0');
INSERT INTO `pz_node` VALUES ('6', '1', '1asd', '', ',1,6,', 'asddddd', '0');
INSERT INTO `pz_node` VALUES ('7', '1', '1asd', '', ',1,7,', '', '0');
INSERT INTO `pz_node` VALUES ('8', '1', 'asd', 'asdasdasd', ',1,8,', '', '0');
INSERT INTO `pz_node` VALUES ('9', '3', '测试1', '测试1测试', ',1,3,9,', '', '0');

-- ----------------------------
-- Table structure for pz_user
-- ----------------------------
DROP TABLE IF EXISTS `pz_user`;
CREATE TABLE `pz_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `state` int(255) NOT NULL DEFAULT '0' COMMENT '状态',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_pz_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_user
-- ----------------------------
INSERT INTO `pz_user` VALUES ('1', 'root', '10470c3b4b1fed12c3baac014be15fac67c6e815', '0', null, null, null);
INSERT INTO `pz_user` VALUES ('2', 'asd', '', '0', null, null, null);
INSERT INTO `pz_user` VALUES ('4', 'asd', '455454', '0', null, null, null);
INSERT INTO `pz_user` VALUES ('5', 'asd', '455454', '0', null, null, null);
INSERT INTO `pz_user` VALUES ('8', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('9', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('10', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('11', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('12', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('13', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('14', 'asd', '312312', '1', null, null, null);
INSERT INTO `pz_user` VALUES ('15', 'asd', '312312', '1', null, null, null);
