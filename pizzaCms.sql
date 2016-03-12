/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50173
Source Host           : 192.168.1.117:3306
Source Database       : pizzaCms

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2016-03-12 19:43:21
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
  `createtime` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `page` (`id`,`title`,`nodeid`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_article
-- ----------------------------
INSERT INTO `pz_article` VALUES ('1', '丈夫将老婆名写篮球上 一生气就打被判定家暴', '/upload/2016/03/12/_3sw2_2acltjopcrqv5brhmhxlzst7wl.jpg', '<div class=\"otitle\" style=\"padding:0px;margin:20px 0px 0px;font-size:14px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	（原标题：他把老婆名字写在篮球上 拍球时不停地说“打死你”）\n</div>\n<div id=\"endText\" class=\"end-text\" style=\"padding:0px 0px 20px;margin:0px 10px 0px 0px;text-align:justify;font-size:16px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	<p style=\"text-indent:2em;\">\n		3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		3月10日上午，是反家暴法生效的第十天，区妇联联合区委政法委、区司法局、区公安局，开展了《反家庭暴力法》业务知识培训。参加会议的有全区妇女代表以及司法局、公安局等相关科室人员，共计200余人参加。\n	</p>\n	<p style=\"text-indent:2em;\">\n		培训会邀请了重庆市经管学院心理学教授、全国公安系统优秀教师郭子贤教授。会上，郭教授用简洁易懂的方式，给大家诠释了反家庭暴力的相关条款。“不孝子女殴打父母，或者妻子殴打丈夫，这些也是家庭暴力。”郭教授说，只要是发生在家庭成员之间的侵害行为，都属于家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		“同居之间的恋人，一方殴打另一方，也是家庭暴力。”郭教授介绍，如今只要是具有共同生活关系，比如同居、扶养、寄养等，他们之间出现的殴打、谩骂，都能算作家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		而人们很少意识到的恐吓，也是家庭暴力的一种。郭教授说，在他接触过的案例中，曾有一个丈夫，因为对妻子不满。便在家中放置了很多篮球，篮球上写上妻子的名字。每天闲来无事，他便拍打篮球，同时口中念念有词“×××，打死你！”等等。\n	</p>\n	<p style=\"text-indent:2em;\">\n		时间一长，妻子的精神受到了极大的伤害，以至于她一听到“篮球”二字就会浑身发抖，要是听到打篮球的声音，就会抱头躲开。最后，经过调查，判定丈夫的这种行为已经构成了家庭暴力。\n	</p>\n</div>', '3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。', '12', '0', '0', '1', '1', '网易新闻', '家暴 反家庭暴力法', 'http://www.baidu.com', '0', '0', '1457779085');
INSERT INTO `pz_article` VALUES ('11', '南非少年发现疑似马航MH370航班客机残片', '', '<p style=\"font-size:16px;text-indent:2em;color:#252525;font-family:宋体, sans-serif;text-align:justify;background-color:#FFFFFF;\">\n	新华社约翰内斯堡3月11日电 据南非媒体11日报道，一名南非少年去年年底在莫桑比克海滩度假时发现疑似马来西亚航空公司MH370航班客机的残片，这块残片将由南非民用航空管理局送往澳大利亚接受鉴定。\n</p>\n<p style=\"font-size:16px;text-indent:2em;color:#252525;font-family:宋体, sans-serif;text-align:justify;background-color:#FFFFFF;\">\n	据报道，去年12月30日，南非少年利亚姆·洛特在莫桑比克南部赛赛地区海滩度假时发现一块长约一米、带铆钉孔的金属片，金属片上还印有“676EB”字样。洛特认为这是飞机残片，因此在度假结束后将金属片带回南非。\n</p>\n<p style=\"font-size:16px;text-indent:2em;color:#252525;font-family:宋体, sans-serif;text-align:justify;background-color:#FFFFFF;\">\n	洛特说，在得知有人在莫桑比克海岸附近发现疑似MH370航班客机残片后，他决定向南非民用航空管理局报告自己的有关发现。\n</p>\n<p style=\"font-size:16px;text-indent:2em;color:#252525;font-family:宋体, sans-serif;text-align:justify;background-color:#FFFFFF;\">\n	南非民用航空管理局表示，洛特发现的这块碎片可能来自一架波音777客机，民用航空管理局将尽快把这块碎片转交给澳大利亚相关机构进行调查。\n</p>\n<p style=\"font-size:16px;text-indent:2em;color:#252525;font-family:宋体, sans-serif;text-align:justify;background-color:#FFFFFF;\">\n	2014年3月8日，从马来西亚吉隆坡飞往中国北京的马来西亚航空公司MH370航班客机失踪，机上载有239人。2015年1月29日，马来西亚民航局宣布该航班客机失事，同时推定机上所有人员遇难。\n</p>', '新华社约翰内斯堡3月11日电 据南非媒体11日报道，一名南非少年去年年底在莫桑比克海滩度假时发现疑似马来西亚航空公司MH370航班客机的残片，这块残片将由南非民用航空管理局送往澳大利亚接受鉴定。', '3', '0', '0', '1', '1', '网易新闻', '', 'baidu.com', '0', '0', '1457779085');

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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_node
-- ----------------------------
INSERT INTO `pz_node` VALUES ('1', '0', '首页', '', ',1,', '', '0');
INSERT INTO `pz_node` VALUES ('3', '1', '国际', '国际豆腐干豆腐干', ',1,3,', '', '0');
INSERT INTO `pz_node` VALUES ('4', '1', '排行', '', ',1,4,', '', '0');
INSERT INTO `pz_node` VALUES ('5', '1', '图片', '', ',1,5,', '', '0');
INSERT INTO `pz_node` VALUES ('6', '1', '国内', '', ',1,6,', '', '0');
INSERT INTO `pz_node` VALUES ('7', '1', '社会', '', ',1,7,', '', '0');
INSERT INTO `pz_node` VALUES ('8', '1', '聚合', '网易聚合阅读', ',1,8,', '', '1');
INSERT INTO `pz_node` VALUES ('9', '3', '国际评论', '测试1测试', ',1,3,9,', '', '0');
INSERT INTO `pz_node` VALUES ('10', '1', '数读', '', ',1,10,', '', '0');
INSERT INTO `pz_node` VALUES ('11', '8', '聚合军事', '', ',1,8,11,', '', '0');
INSERT INTO `pz_node` VALUES ('12', '11', '两会观点', '', ',1,8,11,12,', '', '0');

-- ----------------------------
-- Table structure for pz_user
-- ----------------------------
DROP TABLE IF EXISTS `pz_user`;
CREATE TABLE `pz_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL DEFAULT '',
  `nickname` varchar(30) DEFAULT '' COMMENT '昵称',
  `password` varchar(100) NOT NULL DEFAULT '',
  `state` int(255) NOT NULL DEFAULT '0' COMMENT '状态',
  `salt` varchar(10) NOT NULL DEFAULT 'dx#$59',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_user
-- ----------------------------
INSERT INTO `pz_user` VALUES ('1', 'root', '左盐', 'ad59ca8184ffbcda9953a036ef28c8ad', '0', 'zOBgZ');
INSERT INTO `pz_user` VALUES ('16', 'root1', 'root', '242524fb01200c1cc31f1a6121788fb8', '0', 'spDdr');
INSERT INTO `pz_user` VALUES ('17', 'root1', 'root', 'df8de797279a2f604aa7d3b709635526', '0', 'yEuGz');
INSERT INTO `pz_user` VALUES ('18', 'root1', 'root', '4b1ed106d8c506e13b66feb2e36410cf', '0', 'eN@ul');
