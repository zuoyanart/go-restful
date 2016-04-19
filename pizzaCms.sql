/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50173
Source Host           : 192.168.1.117:3306
Source Database       : pizzaCms

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2016-04-19 20:56:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pz_action
-- ----------------------------
DROP TABLE IF EXISTS `pz_action`;
CREATE TABLE `pz_action` (
  `actionId` int(11) NOT NULL AUTO_INCREMENT,
  `actionName` varchar(255) NOT NULL DEFAULT '' COMMENT '动作名称',
  `action` varchar(255) NOT NULL DEFAULT '' COMMENT '动作编码，英文或数字',
  `actionModelId` int(11) DEFAULT '0' COMMENT '栏目模块id',
  `actionState` int(255) DEFAULT '0' COMMENT '动作状态,0有效',
  PRIMARY KEY (`actionId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_action
-- ----------------------------

-- ----------------------------
-- Table structure for pz_actionModel
-- ----------------------------
DROP TABLE IF EXISTS `pz_actionModel`;
CREATE TABLE `pz_actionModel` (
  `actionModelId` int(11) NOT NULL AUTO_INCREMENT,
  `actionModelName` varchar(255) NOT NULL DEFAULT '' COMMENT '栏目模块名称',
  `actionModel` varchar(255) NOT NULL DEFAULT '' COMMENT '栏目模块编码',
  `actionModelState` int(255) DEFAULT '0' COMMENT '栏目模块状态',
  PRIMARY KEY (`actionModelId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_actionModel
-- ----------------------------

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_article
-- ----------------------------
INSERT INTO `pz_article` VALUES ('1', '丈夫将老婆名写篮球上 一生气就打被判定家暴1', '/upload/2016/03/12/_3sw2_2acltjopcrqv5brhmhxlzst7wl.jpg', '<div class=\"otitle\" style=\"padding:0px;margin:20px 0px 0px;font-size:14px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	（原标题：他把老婆名字写在篮球上 拍球时不停地说“打死你”）\n</div>\n<div id=\"endText\" class=\"end-text\" style=\"padding:0px 0px 20px;margin:0px 10px 0px 0px;text-align:justify;font-size:16px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	<p style=\"text-indent:2em;\">\n		3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		3月10日上午，是反家暴法生效的第十天，区妇联联合区委政法委、区司法局、区公安局，开展了《反家庭暴力法》业务知识培训。参加会议的有全区妇女代表以及司法局、公安局等相关科室人员，共计200余人参加。\n	</p>\n	<p style=\"text-indent:2em;\">\n		培训会邀请了重庆市经管学院心理学教授、全国公安系统优秀教师郭子贤教授。会上，郭教授用简洁易懂的方式，给大家诠释了反家庭暴力的相关条款。“不孝子女殴打父母，或者妻子殴打丈夫，这些也是家庭暴力。”郭教授说，只要是发生在家庭成员之间的侵害行为，都属于家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		“同居之间的恋人，一方殴打另一方，也是家庭暴力。”郭教授介绍，如今只要是具有共同生活关系，比如同居、扶养、寄养等，他们之间出现的殴打、谩骂，都能算作家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		而人们很少意识到的恐吓，也是家庭暴力的一种。郭教授说，在他接触过的案例中，曾有一个丈夫，因为对妻子不满。便在家中放置了很多篮球，篮球上写上妻子的名字。每天闲来无事，他便拍打篮球，同时口中念念有词“×××，打死你！”等等。\n	</p>\n	<p style=\"text-indent:2em;\">\n		时间一长，妻子的精神受到了极大的伤害，以至于她一听到“篮球”二字就会浑身发抖，要是听到打篮球的声音，就会抱头躲开。最后，经过调查，判定丈夫的这种行为已经构成了家庭暴力。\n	</p>\n</div>', '3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。', '8', '0', '0', '1', '1', '网易新闻', '家暴 反家庭暴力法', 'http://www.baidu.com', '0', '0', '1457779085');
INSERT INTO `pz_article` VALUES ('2', '丈夫将老婆名写篮球上 一生气就打被判定家暴2', '/upload/2016/03/12/_3sw2_2acltjopcrqv5brhmhxlzst7wl.jpg', '<div class=\"otitle\" style=\"padding:0px;margin:20px 0px 0px;font-size:14px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	（原标题：他把老婆名字写在篮球上 拍球时不停地说“打死你”）\n</div>\n<div id=\"endText\" class=\"end-text\" style=\"padding:0px 0px 20px;margin:0px 10px 0px 0px;text-align:justify;font-size:16px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	<p style=\"text-indent:2em;\">\n		3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		3月10日上午，是反家暴法生效的第十天，区妇联联合区委政法委、区司法局、区公安局，开展了《反家庭暴力法》业务知识培训。参加会议的有全区妇女代表以及司法局、公安局等相关科室人员，共计200余人参加。\n	</p>\n	<p style=\"text-indent:2em;\">\n		培训会邀请了重庆市经管学院心理学教授、全国公安系统优秀教师郭子贤教授。会上，郭教授用简洁易懂的方式，给大家诠释了反家庭暴力的相关条款。“不孝子女殴打父母，或者妻子殴打丈夫，这些也是家庭暴力。”郭教授说，只要是发生在家庭成员之间的侵害行为，都属于家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		“同居之间的恋人，一方殴打另一方，也是家庭暴力。”郭教授介绍，如今只要是具有共同生活关系，比如同居、扶养、寄养等，他们之间出现的殴打、谩骂，都能算作家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		而人们很少意识到的恐吓，也是家庭暴力的一种。郭教授说，在他接触过的案例中，曾有一个丈夫，因为对妻子不满。便在家中放置了很多篮球，篮球上写上妻子的名字。每天闲来无事，他便拍打篮球，同时口中念念有词“×××，打死你！”等等。\n	</p>\n	<p style=\"text-indent:2em;\">\n		时间一长，妻子的精神受到了极大的伤害，以至于她一听到“篮球”二字就会浑身发抖，要是听到打篮球的声音，就会抱头躲开。最后，经过调查，判定丈夫的这种行为已经构成了家庭暴力。\n	</p>\n</div>', '3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。', '12', '0', '0', '1', '1', '网易新闻', '家暴 反家庭暴力法', 'http://www.baidu.com', '0', '0', '1457779085');
INSERT INTO `pz_article` VALUES ('3', '丈夫将老婆名写篮球上 一生气就打被判定家暴3', '/upload/2016/03/12/_3sw2_2acltjopcrqv5brhmhxlzst7wl.jpg', '<div class=\"otitle\" style=\"padding:0px;margin:20px 0px 0px;font-size:14px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	（原标题：他把老婆名字写在篮球上 拍球时不停地说“打死你”）\n</div>\n<div id=\"endText\" class=\"end-text\" style=\"padding:0px 0px 20px;margin:0px 10px 0px 0px;text-align:justify;font-size:16px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	<p style=\"text-indent:2em;\">\n		3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		3月10日上午，是反家暴法生效的第十天，区妇联联合区委政法委、区司法局、区公安局，开展了《反家庭暴力法》业务知识培训。参加会议的有全区妇女代表以及司法局、公安局等相关科室人员，共计200余人参加。\n	</p>\n	<p style=\"text-indent:2em;\">\n		培训会邀请了重庆市经管学院心理学教授、全国公安系统优秀教师郭子贤教授。会上，郭教授用简洁易懂的方式，给大家诠释了反家庭暴力的相关条款。“不孝子女殴打父母，或者妻子殴打丈夫，这些也是家庭暴力。”郭教授说，只要是发生在家庭成员之间的侵害行为，都属于家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		“同居之间的恋人，一方殴打另一方，也是家庭暴力。”郭教授介绍，如今只要是具有共同生活关系，比如同居、扶养、寄养等，他们之间出现的殴打、谩骂，都能算作家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		而人们很少意识到的恐吓，也是家庭暴力的一种。郭教授说，在他接触过的案例中，曾有一个丈夫，因为对妻子不满。便在家中放置了很多篮球，篮球上写上妻子的名字。每天闲来无事，他便拍打篮球，同时口中念念有词“×××，打死你！”等等。\n	</p>\n	<p style=\"text-indent:2em;\">\n		时间一长，妻子的精神受到了极大的伤害，以至于她一听到“篮球”二字就会浑身发抖，要是听到打篮球的声音，就会抱头躲开。最后，经过调查，判定丈夫的这种行为已经构成了家庭暴力。\n	</p>\n</div>', '3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。', '10', '0', '0', '1', '1', '网易新闻', '家暴 反家庭暴力法', 'http://www.baidu.com', '0', '0', '1457779085');
INSERT INTO `pz_article` VALUES ('4', '丈夫将老婆名写篮球上 一生气就打被判定家暴4', '/upload/2016/03/12/_3sw2_2acltjopcrqv5brhmhxlzst7wl.jpg', '<div class=\"otitle\" style=\"padding:0px;margin:20px 0px 0px;font-size:14px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	（原标题：他把老婆名字写在篮球上 拍球时不停地说“打死你”）\n</div>\n<div id=\"endText\" class=\"end-text\" style=\"padding:0px 0px 20px;margin:0px 10px 0px 0px;text-align:justify;font-size:16px;color:#252525;font-family:宋体, sans-serif;background-color:#FFFFFF;\">\n	<p style=\"text-indent:2em;\">\n		3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		3月10日上午，是反家暴法生效的第十天，区妇联联合区委政法委、区司法局、区公安局，开展了《反家庭暴力法》业务知识培训。参加会议的有全区妇女代表以及司法局、公安局等相关科室人员，共计200余人参加。\n	</p>\n	<p style=\"text-indent:2em;\">\n		培训会邀请了重庆市经管学院心理学教授、全国公安系统优秀教师郭子贤教授。会上，郭教授用简洁易懂的方式，给大家诠释了反家庭暴力的相关条款。“不孝子女殴打父母，或者妻子殴打丈夫，这些也是家庭暴力。”郭教授说，只要是发生在家庭成员之间的侵害行为，都属于家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		“同居之间的恋人，一方殴打另一方，也是家庭暴力。”郭教授介绍，如今只要是具有共同生活关系，比如同居、扶养、寄养等，他们之间出现的殴打、谩骂，都能算作家庭暴力。\n	</p>\n	<p style=\"text-indent:2em;\">\n		而人们很少意识到的恐吓，也是家庭暴力的一种。郭教授说，在他接触过的案例中，曾有一个丈夫，因为对妻子不满。便在家中放置了很多篮球，篮球上写上妻子的名字。每天闲来无事，他便拍打篮球，同时口中念念有词“×××，打死你！”等等。\n	</p>\n	<p style=\"text-indent:2em;\">\n		时间一长，妻子的精神受到了极大的伤害，以至于她一听到“篮球”二字就会浑身发抖，要是听到打篮球的声音，就会抱头躲开。最后，经过调查，判定丈夫的这种行为已经构成了家庭暴力。\n	</p>\n</div>', '3月1日，我国第一部《反家庭暴力法》正式实施，意味着家庭暴力属于“家务事”的时代正式终结。除了大家都清楚的，家庭成员之间的侵害行为，属于家庭暴力。反家暴法还适用于具有共同生活关系的成员，也就是说，情侣同居出现殴打、谩骂等行为，也是家庭暴力。', '7', '0', '0', '1', '1', '网易新闻', '家暴 反家庭暴力法', 'http://www.baidu.com', '0', '0', '1457779085');

-- ----------------------------
-- Table structure for pz_block
-- ----------------------------
DROP TABLE IF EXISTS `pz_block`;
CREATE TABLE `pz_block` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '',
  `content` varchar(10000) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=106 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_block
-- ----------------------------
INSERT INTO `pz_block` VALUES ('2', 'title', 'content');
INSERT INTO `pz_block` VALUES ('8', 'title', 'content');
INSERT INTO `pz_block` VALUES ('3', 'title', 'content');
INSERT INTO `pz_block` VALUES ('4', 'title', 'content');
INSERT INTO `pz_block` VALUES ('5', 'title', 'content');
INSERT INTO `pz_block` VALUES ('6', 'title', 'content');
INSERT INTO `pz_block` VALUES ('7', 'title', 'content');
INSERT INTO `pz_block` VALUES ('48', 'title', 'content');
INSERT INTO `pz_block` VALUES ('47', 'title', 'content');
INSERT INTO `pz_block` VALUES ('49', 'title', 'content');
INSERT INTO `pz_block` VALUES ('50', 'title', 'content');
INSERT INTO `pz_block` VALUES ('51', 'title', 'content');
INSERT INTO `pz_block` VALUES ('52', 'title', 'content');
INSERT INTO `pz_block` VALUES ('53', 'title', 'content');
INSERT INTO `pz_block` VALUES ('54', 'title', 'content');
INSERT INTO `pz_block` VALUES ('55', 'title', 'content');
INSERT INTO `pz_block` VALUES ('56', 'title', 'content');
INSERT INTO `pz_block` VALUES ('57', 'title', 'content');

-- ----------------------------
-- Table structure for pz_comment
-- ----------------------------
DROP TABLE IF EXISTS `pz_comment`;
CREATE TABLE `pz_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `articleid` int(11) DEFAULT '0' COMMENT '文章id',
  `addtime` int(11) DEFAULT '0' COMMENT '添加时间',
  `content` varchar(1000) DEFAULT '' COMMENT '评论内容',
  `uid` int(11) DEFAULT '0' COMMENT '用户id',
  `username` varchar(30) DEFAULT '' COMMENT '用户昵称',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=254 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_comment
-- ----------------------------
INSERT INTO `pz_comment` VALUES ('17', '2', '1', 'content17', '1', 'username');
INSERT INTO `pz_comment` VALUES ('18', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('19', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('20', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('21', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('22', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('23', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('24', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('25', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('26', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('27', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('28', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('29', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('30', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('31', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('32', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('33', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('34', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('35', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('36', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('37', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('38', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('39', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('40', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('41', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('42', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('43', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('44', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('45', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('46', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('47', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('48', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('49', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('50', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('51', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('52', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('53', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('54', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('55', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('56', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('57', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('58', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('59', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('60', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('61', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('62', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('63', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('64', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('65', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('66', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('67', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('68', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('69', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('70', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('71', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('72', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('73', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('74', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('75', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('76', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('77', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('78', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('79', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('80', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('81', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('82', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('83', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('84', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('85', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('86', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('87', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('88', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('89', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('90', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('91', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('92', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('93', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('94', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('95', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('96', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('97', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('98', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('99', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('100', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('101', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('102', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('103', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('104', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('105', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('106', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('107', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('108', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('109', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('110', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('111', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('112', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('113', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('114', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('115', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('116', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('117', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('118', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('119', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('120', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('121', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('122', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('123', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('124', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('125', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('126', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('127', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('128', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('129', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('130', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('131', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('132', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('133', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('134', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('135', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('136', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('137', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('138', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('139', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('140', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('141', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('142', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('143', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('144', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('145', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('146', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('147', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('148', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('149', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('150', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('151', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('152', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('153', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('154', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('155', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('156', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('157', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('158', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('159', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('160', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('161', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('162', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('163', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('164', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('165', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('166', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('167', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('168', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('169', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('170', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('171', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('172', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('173', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('174', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('175', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('176', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('177', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('178', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('179', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('180', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('181', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('182', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('183', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('184', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('185', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('186', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('187', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('188', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('189', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('190', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('191', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('192', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('193', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('194', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('195', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('196', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('197', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('198', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('199', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('200', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('201', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('202', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('203', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('204', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('205', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('206', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('207', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('208', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('209', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('210', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('211', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('212', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('213', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('214', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('215', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('216', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('217', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('218', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('219', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('220', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('221', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('222', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('223', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('224', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('225', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('226', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('227', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('228', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('229', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('230', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('231', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('232', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('233', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('234', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('235', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('236', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('237', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('238', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('239', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('240', '2', '1', 'content', '1', 'username');
INSERT INTO `pz_comment` VALUES ('241', '2', '1', 'content', '1', 'username');

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
-- Table structure for pz_role
-- ----------------------------
DROP TABLE IF EXISTS `pz_role`;
CREATE TABLE `pz_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupid` int(11) DEFAULT '0' COMMENT '用户组id',
  `name` varchar(255) DEFAULT '' COMMENT '角色名称',
  `des` varchar(1000) DEFAULT '' COMMENT '角色描述',
  `state` int(255) DEFAULT '0' COMMENT '角色状态',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_role
-- ----------------------------
INSERT INTO `pz_role` VALUES ('3', '3', '主编', '', '0');
INSERT INTO `pz_role` VALUES ('4', '3', '编辑', '', '0');

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
  `userGroupId` int(11) DEFAULT '0' COMMENT '用户组id',
  `roleId` int(11) DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_user
-- ----------------------------
INSERT INTO `pz_user` VALUES ('1', 'root', '左盐', 'ca961093927fe366611a46e370d89e4f', '0', 'x#WSWi%oGM', '0', '0');

-- ----------------------------
-- Table structure for pz_userGroup
-- ----------------------------
DROP TABLE IF EXISTS `pz_userGroup`;
CREATE TABLE `pz_userGroup` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '' COMMENT '用户组名称',
  `des` varchar(1000) DEFAULT '' COMMENT '用户组描述',
  `state` int(4) DEFAULT '0' COMMENT '用户组状态',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pz_userGroup
-- ----------------------------
INSERT INTO `pz_userGroup` VALUES ('1', '超级管理员', '超级管理员', '0');
INSERT INTO `pz_userGroup` VALUES ('3', '编辑部', '编辑部', '0');
