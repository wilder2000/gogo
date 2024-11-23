/*
 Navicat MySQL Data Transfer

 Source Server         : gomysql
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : localhost:3306
 Source Schema         : gohome

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 21/07/2022 21:39:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for a_clubdocs
-- ----------------------------
DROP TABLE IF EXISTS `a_clubdocs`;
CREATE TABLE `a_clubdocs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `clubid` int(11) DEFAULT NULL COMMENT '团队编号与a_clubs中的id关联',
  `docid` int(11) DEFAULT NULL COMMENT '文档编号，与a_docs中的id关联',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '加入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='团队文档表';

-- ----------------------------
-- Table structure for a_clubs
-- ----------------------------
DROP TABLE IF EXISTS `a_clubs`;
CREATE TABLE `a_clubs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '团队空间名称',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `describe` varchar(255) DEFAULT NULL COMMENT '空间描述',
  `ownerid` varchar(50) DEFAULT NULL COMMENT '用户编号，与s_user关联',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='团队空间表';

-- ----------------------------
-- Table structure for a_docnotes
-- ----------------------------
DROP TABLE IF EXISTS `a_docnotes`;
CREATE TABLE `a_docnotes` (
  `id` int(11) DEFAULT NULL,
  `userid` varchar(50) DEFAULT NULL COMMENT '用户编号s_user中的id',
  `docid` int(11) DEFAULT NULL COMMENT '文档编号',
  `notes` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档批注表';

-- ----------------------------
-- Table structure for a_docs
-- ----------------------------
DROP TABLE IF EXISTS `a_docs`;
CREATE TABLE `a_docs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL COMMENT '文档标题',
  `describe` varchar(255) DEFAULT NULL COMMENT '文档描述',
  `url` varchar(100) DEFAULT NULL COMMENT '存放地址',
  `uploadtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `tag` int(11) DEFAULT NULL COMMENT '标签的编号，与a_tags中的id关联',
  `ownerid` varchar(50) DEFAULT NULL COMMENT '文档拥有者编号',
  `grouprul` int(11) DEFAULT NULL COMMENT '组织内部权限\n7=4（r），2(w)，1(x)，0没有权限',
  `otherrul` int(11) DEFAULT NULL COMMENT '组织外部权限',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档表';

-- ----------------------------
-- Table structure for a_shares
-- ----------------------------
DROP TABLE IF EXISTS `a_shares`;
CREATE TABLE `a_shares` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '分享名称',
  `code` varchar(20) DEFAULT NULL COMMENT '分享链接编码用于接入，生成访问链接',
  `docid` int(11) NOT NULL COMMENT '文档编号',
  `validtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '有效期，过期不能访问',
  `clicktimes` int(11) DEFAULT NULL COMMENT '点击次数',
  `state` int(11) DEFAULT NULL COMMENT '状态：1可以访问，0,禁用',
  `type` int(11) DEFAULT NULL COMMENT '访问类型：1外部公开，不要账号；0,需要使用账号才能访问',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `code_idx` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档分享表';

-- ----------------------------
-- Table structure for a_tags
-- ----------------------------
DROP TABLE IF EXISTS `a_tags`;
CREATE TABLE `a_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL COMMENT '标签名称',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `sumdocs` int(11) DEFAULT NULL COMMENT '文档数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档标签表';

-- ----------------------------
-- Table structure for a_userdocs
-- ----------------------------
DROP TABLE IF EXISTS `a_userdocs`;
CREATE TABLE `a_userdocs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '流水号',
  `userid` varchar(50) DEFAULT NULL COMMENT '用户编号，与s_user中id关联',
  `docid` int(11) NOT NULL COMMENT '文档编号表',
  `owner` int(11) DEFAULT NULL COMMENT '是否文档拥有人',
  `reader` int(11) DEFAULT NULL COMMENT '可以读取文档',
  `modifyer` int(11) DEFAULT NULL COMMENT '可以修改文档',
  `deleter` int(11) DEFAULT NULL COMMENT '是否要删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户的文档表';

-- ----------------------------
-- Table structure for hc_admin
-- ----------------------------
DROP TABLE IF EXISTS `hc_admin`;
CREATE TABLE `hc_admin` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Username` varchar(16) NOT NULL DEFAULT '',
  `Password` varchar(32) NOT NULL DEFAULT '',
  `Email` varchar(50) NOT NULL DEFAULT '',
  `Code` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Username` (`Username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_banner
-- ----------------------------
DROP TABLE IF EXISTS `hc_banner`;
CREATE TABLE `hc_banner` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Picture` varchar(50) NOT NULL DEFAULT '',
  `Title` varchar(255) NOT NULL DEFAULT '',
  `Link` varchar(255) NOT NULL DEFAULT '',
  `Sort` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_category
-- ----------------------------
DROP TABLE IF EXISTS `hc_category`;
CREATE TABLE `hc_category` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Pid` int(11) NOT NULL DEFAULT '0',
  `Title` varchar(20) NOT NULL DEFAULT '',
  `Cover` varchar(255) NOT NULL DEFAULT '',
  `Cnt` int(11) NOT NULL DEFAULT '0',
  `Sort` int(11) NOT NULL DEFAULT '0',
  `Alias` varchar(30) NOT NULL DEFAULT '',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Pid` (`Pid`,`Title`)
) ENGINE=InnoDB AUTO_INCREMENT=315 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_coin_log
-- ----------------------------
DROP TABLE IF EXISTS `hc_coin_log`;
CREATE TABLE `hc_coin_log` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Coin` int(11) NOT NULL DEFAULT '0',
  `Log` varchar(512) NOT NULL DEFAULT '',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `hc_coin_log_Uid` (`Uid`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_collect
-- ----------------------------
DROP TABLE IF EXISTS `hc_collect`;
CREATE TABLE `hc_collect` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Cid` int(11) NOT NULL DEFAULT '0',
  `Did` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Did` (`Did`,`Cid`),
  KEY `hc_collect_Cid` (`Cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_collect_folder
-- ----------------------------
DROP TABLE IF EXISTS `hc_collect_folder`;
CREATE TABLE `hc_collect_folder` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Cover` varchar(50) NOT NULL DEFAULT '',
  `Title` varchar(100) NOT NULL DEFAULT '默认收藏夹',
  `Description` varchar(512) NOT NULL DEFAULT '',
  `Uid` int(11) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `Cnt` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Title` (`Title`,`Uid`),
  KEY `hc_collect_folder_Uid` (`Uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_config
-- ----------------------------
DROP TABLE IF EXISTS `hc_config`;
CREATE TABLE `hc_config` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) NOT NULL DEFAULT '',
  `InputType` varchar(10) NOT NULL DEFAULT '',
  `Description` varchar(255) NOT NULL DEFAULT '',
  `Key` varchar(30) NOT NULL DEFAULT '',
  `Value` varchar(255) NOT NULL DEFAULT '',
  `Category` varchar(30) NOT NULL DEFAULT '',
  `Options` varchar(4096) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Key` (`Key`,`Category`),
  KEY `hc_config_Category` (`Category`)
) ENGINE=InnoDB AUTO_INCREMENT=2961 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_doc_text
-- ----------------------------
DROP TABLE IF EXISTS `hc_doc_text`;
CREATE TABLE `hc_doc_text` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Md5` varchar(32) NOT NULL DEFAULT '',
  `Content` varchar(5000) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document
-- ----------------------------
DROP TABLE IF EXISTS `hc_document`;
CREATE TABLE `hc_document` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) NOT NULL DEFAULT '',
  `Filename` varchar(255) NOT NULL DEFAULT '',
  `Keywords` varchar(255) NOT NULL DEFAULT '',
  `Description` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_comment
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_comment`;
CREATE TABLE `hc_document_comment` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Did` int(11) NOT NULL DEFAULT '0',
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Score` int(11) NOT NULL DEFAULT '30000',
  `Content` varchar(256) NOT NULL DEFAULT '',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Did` (`Did`,`Uid`),
  KEY `hc_document_comment_Did` (`Did`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_illegal
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_illegal`;
CREATE TABLE `hc_document_illegal` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Md5` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Md5` (`Md5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_info
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_info`;
CREATE TABLE `hc_document_info` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `DsId` int(11) NOT NULL DEFAULT '0',
  `Uid` int(11) NOT NULL DEFAULT '0',
  `ChanelId` int(11) NOT NULL DEFAULT '0',
  `Pid` int(11) NOT NULL DEFAULT '0',
  `Cid` int(11) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `TimeUpdate` int(11) NOT NULL DEFAULT '0',
  `Dcnt` int(11) NOT NULL DEFAULT '0',
  `Vcnt` int(11) NOT NULL DEFAULT '0',
  `Ccnt` int(11) NOT NULL DEFAULT '0',
  `Score` int(11) NOT NULL DEFAULT '30000',
  `ScorePeople` int(11) NOT NULL DEFAULT '0',
  `Price` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `hc_document_info_DsId` (`DsId`),
  KEY `hc_document_info_Uid` (`Uid`),
  KEY `hc_document_info_ChanelId` (`ChanelId`),
  KEY `hc_document_info_Pid` (`Pid`),
  KEY `hc_document_info_Cid` (`Cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_recycle
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_recycle`;
CREATE TABLE `hc_document_recycle` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Date` int(11) NOT NULL DEFAULT '0',
  `Self` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_remark
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_remark`;
CREATE TABLE `hc_document_remark` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Content` varchar(5120) NOT NULL DEFAULT '',
  `AllowDownload` tinyint(1) NOT NULL DEFAULT '1',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_document_store
-- ----------------------------
DROP TABLE IF EXISTS `hc_document_store`;
CREATE TABLE `hc_document_store` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Md5` varchar(32) NOT NULL DEFAULT '',
  `Ext` varchar(10) NOT NULL DEFAULT '',
  `ExtCate` varchar(10) NOT NULL DEFAULT '',
  `ExtNum` int(11) NOT NULL DEFAULT '0',
  `Page` int(11) NOT NULL DEFAULT '0',
  `PreviewPage` int(11) NOT NULL DEFAULT '50',
  `Size` int(11) NOT NULL DEFAULT '0',
  `ModTime` int(11) NOT NULL DEFAULT '0',
  `PreviewExt` varchar(4) NOT NULL DEFAULT 'svg',
  `Width` int(11) NOT NULL DEFAULT '0',
  `Height` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Md5` (`Md5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_free_down
-- ----------------------------
DROP TABLE IF EXISTS `hc_free_down`;
CREATE TABLE `hc_free_down` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Did` int(11) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_friend
-- ----------------------------
DROP TABLE IF EXISTS `hc_friend`;
CREATE TABLE `hc_friend` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(100) NOT NULL DEFAULT '',
  `Link` varchar(100) NOT NULL DEFAULT '',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  `Sort` int(11) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Link` (`Link`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_pages
-- ----------------------------
DROP TABLE IF EXISTS `hc_pages`;
CREATE TABLE `hc_pages` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) NOT NULL DEFAULT '',
  `Alias` varchar(30) NOT NULL DEFAULT '',
  `Title` varchar(255) NOT NULL DEFAULT '',
  `Keywords` varchar(255) NOT NULL DEFAULT '',
  `Description` varchar(255) NOT NULL DEFAULT '',
  `Content` varchar(5120) NOT NULL DEFAULT '',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `Sort` int(11) NOT NULL DEFAULT '100',
  `Vcnt` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Alias` (`Alias`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_relate
-- ----------------------------
DROP TABLE IF EXISTS `hc_relate`;
CREATE TABLE `hc_relate` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Ids` varchar(512) NOT NULL DEFAULT '',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_report
-- ----------------------------
DROP TABLE IF EXISTS `hc_report`;
CREATE TABLE `hc_report` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Did` int(11) NOT NULL DEFAULT '0',
  `Reason` int(11) NOT NULL DEFAULT '1',
  `Status` tinyint(1) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `TimeUpdate` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Uid` (`Uid`,`Did`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_search_log
-- ----------------------------
DROP TABLE IF EXISTS `hc_search_log`;
CREATE TABLE `hc_search_log` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Wd` varchar(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Wd` (`Wd`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_seo
-- ----------------------------
DROP TABLE IF EXISTS `hc_seo`;
CREATE TABLE `hc_seo` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL DEFAULT '',
  `Page` varchar(30) NOT NULL DEFAULT '',
  `IsMobile` tinyint(1) NOT NULL DEFAULT '0',
  `Title` varchar(255) NOT NULL DEFAULT '{title}',
  `Keywords` varchar(255) NOT NULL DEFAULT '{keywords}',
  `Description` varchar(255) NOT NULL DEFAULT '{description}',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Page` (`Page`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_sign
-- ----------------------------
DROP TABLE IF EXISTS `hc_sign`;
CREATE TABLE `hc_sign` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Date` varchar(8) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Uid` (`Uid`,`Date`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_suggest
-- ----------------------------
DROP TABLE IF EXISTS `hc_suggest`;
CREATE TABLE `hc_suggest` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) NOT NULL DEFAULT '0',
  `Content` varchar(512) NOT NULL DEFAULT '',
  `Email` varchar(50) NOT NULL DEFAULT '',
  `Name` varchar(20) NOT NULL DEFAULT '',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `TimeUpdate` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_sys
-- ----------------------------
DROP TABLE IF EXISTS `hc_sys`;
CREATE TABLE `hc_sys` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Site` varchar(100) NOT NULL DEFAULT '',
  `Tongji` varchar(2048) NOT NULL DEFAULT '',
  `CntDoc` int(11) NOT NULL DEFAULT '0',
  `CntUser` int(11) NOT NULL DEFAULT '0',
  `Reward` int(11) NOT NULL DEFAULT '5',
  `MaxFile` int(11) NOT NULL DEFAULT '52428800',
  `Sign` int(11) NOT NULL DEFAULT '1',
  `ListRows` int(11) NOT NULL DEFAULT '10',
  `Icp` varchar(255) NOT NULL DEFAULT '',
  `DirtyWord` varchar(2048) NOT NULL DEFAULT '',
  `TimeExpireRelate` int(11) NOT NULL DEFAULT '604800',
  `TimeExpireHotspot` int(11) NOT NULL DEFAULT '604800',
  `MobileOn` tinyint(1) NOT NULL DEFAULT '1',
  `TplMobile` varchar(255) NOT NULL DEFAULT 'default',
  `TplComputer` varchar(255) NOT NULL DEFAULT 'default',
  `TplAdmin` varchar(255) NOT NULL DEFAULT 'default',
  `TplEmailReg` varchar(2048) NOT NULL DEFAULT '',
  `TplEmailFindPwd` varchar(2048) NOT NULL DEFAULT '',
  `DomainPc` varchar(100) NOT NULL DEFAULT 'dochub.me',
  `DomainMobile` varchar(100) NOT NULL DEFAULT 'm.dochub.me',
  `PreviewPage` int(11) NOT NULL DEFAULT '50',
  `Trends` varchar(255) NOT NULL DEFAULT '',
  `FreeDay` int(11) NOT NULL DEFAULT '7',
  `Question` varchar(255) NOT NULL DEFAULT 'DocHub文库的中文名是？',
  `Answer` varchar(255) NOT NULL DEFAULT '多哈',
  `CoinReg` int(11) NOT NULL DEFAULT '10',
  `Watermark` varchar(255) NOT NULL DEFAULT '',
  `ReportReasons` varchar(2048) NOT NULL DEFAULT '',
  `IsCloseReg` tinyint(1) NOT NULL DEFAULT '0',
  `StoreType` varchar(15) NOT NULL DEFAULT 'cs-oss',
  `CheckRegEmail` tinyint(1) NOT NULL DEFAULT '1',
  `AllowRepeatedDoc` tinyint(1) NOT NULL DEFAULT '0',
  `AutoSitemap` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_user
-- ----------------------------
DROP TABLE IF EXISTS `hc_user`;
CREATE TABLE `hc_user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Email` varchar(50) NOT NULL DEFAULT '',
  `Password` varchar(32) NOT NULL DEFAULT '',
  `Username` varchar(16) NOT NULL DEFAULT '',
  `Intro` varchar(255) NOT NULL DEFAULT '',
  `Avatar` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Email` (`Email`),
  UNIQUE KEY `Username` (`Username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_user_info
-- ----------------------------
DROP TABLE IF EXISTS `hc_user_info`;
CREATE TABLE `hc_user_info` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Coin` int(11) NOT NULL DEFAULT '10',
  `Document` int(11) NOT NULL DEFAULT '0',
  `Collect` int(11) NOT NULL DEFAULT '0',
  `TimeCreate` int(11) NOT NULL DEFAULT '0',
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`),
  KEY `hc_user_info_Coin` (`Coin`),
  KEY `hc_user_info_Document` (`Document`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for hc_word
-- ----------------------------
DROP TABLE IF EXISTS `hc_word`;
CREATE TABLE `hc_word` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Wd` varchar(20) NOT NULL DEFAULT '',
  `Cnt` int(11) NOT NULL DEFAULT '0',
  `Ids` longtext NOT NULL,
  `Status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Wd` (`Wd`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for s_departments
-- ----------------------------
DROP TABLE IF EXISTS `s_departments`;
CREATE TABLE `s_departments` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) DEFAULT NULL COMMENT '组织名称',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='组织表';

-- ----------------------------
-- Table structure for s_depusers
-- ----------------------------
DROP TABLE IF EXISTS `s_depusers`;
CREATE TABLE `s_depusers` (
  `userid` varchar(50) NOT NULL COMMENT '用户编号',
  `departmentid` int(11) NOT NULL COMMENT '组织编号',
  PRIMARY KEY (`userid`,`departmentid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='组织成员表';

-- ----------------------------
-- Table structure for s_logs
-- ----------------------------
DROP TABLE IF EXISTS `s_logs`;
CREATE TABLE `s_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) DEFAULT NULL,
  `ip` varchar(50) DEFAULT NULL,
  `logintime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for s_users
-- ----------------------------
DROP TABLE IF EXISTS `s_users`;
CREATE TABLE `s_users` (
  `id` varchar(50) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `icon` varchar(100) DEFAULT NULL COMMENT '头像的url',
  `state` int(11) DEFAULT NULL COMMENT '状态：1，可用，0禁用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
