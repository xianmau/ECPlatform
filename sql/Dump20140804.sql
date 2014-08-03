CREATE DATABASE  IF NOT EXISTS `ecplatformdb` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ecplatformdb`;
-- MySQL dump 10.13  Distrib 5.6.17, for Win32 (x86)
--
-- Host: 127.0.0.1    Database: ecplatformdb
-- ------------------------------------------------------
-- Server version	5.6.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_admin`
--

DROP TABLE IF EXISTS `tb_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_admin` (
  `Name` varchar(128) NOT NULL,
  `Password` varchar(128) DEFAULT NULL,
  `Role` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_admin`
--

LOCK TABLES `tb_admin` WRITE;
/*!40000 ALTER TABLE `tb_admin` DISABLE KEYS */;
INSERT INTO `tb_admin` VALUES ('admin','admin','管理员'),('root','root','root'),('xianmau','123456','管理员'),('xm','123456','管理员');
/*!40000 ALTER TABLE `tb_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_article`
--

DROP TABLE IF EXISTS `tb_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_article` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(128) NOT NULL,
  `Category` varchar(128) NOT NULL,
  `Content` text NOT NULL,
  `Time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Click` int(11) NOT NULL DEFAULT '0',
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_article`
--

LOCK TABLES `tb_article` WRITE;
/*!40000 ALTER TABLE `tb_article` DISABLE KEYS */;
INSERT INTO `tb_article` VALUES (1,'一村一品平台公测','新闻资讯','123','2014-07-20 15:41:27',0,0);
/*!40000 ALTER TABLE `tb_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_article_category`
--

DROP TABLE IF EXISTS `tb_article_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_article_category` (
  `Name` varchar(128) NOT NULL,
  `Parent` varchar(128) NOT NULL,
  `Ordering` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_article_category`
--

LOCK TABLES `tb_article_category` WRITE;
/*!40000 ALTER TABLE `tb_article_category` DISABLE KEYS */;
INSERT INTO `tb_article_category` VALUES ('商城公告','',254),('帮助中心','',253),('新闻资讯','',255),('服务承诺','',252);
/*!40000 ALTER TABLE `tb_article_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_goods`
--

DROP TABLE IF EXISTS `tb_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_goods` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(128) NOT NULL,
  `Category` varchar(128) NOT NULL,
  `Content` text NOT NULL,
  `Origin` varchar(128) NOT NULL,
  `Unit` varchar(128) NOT NULL,
  `Price` decimal(8,2) NOT NULL DEFAULT '0.00',
  `Shop` varchar(128) NOT NULL,
  `BuyLink` varchar(128) NOT NULL,
  `Images` text NOT NULL,
  `Certificates` text NOT NULL,
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods`
--

LOCK TABLES `tb_goods` WRITE;
/*!40000 ALTER TABLE `tb_goods` DISABLE KEYS */;
INSERT INTO `tb_goods` VALUES (1,'华南丝苗米','米','<p>华南丝苗米</p>\n','广东省韶关市乐昌县庆云镇','5KG',138.00,'广州稻草人','http://gzdcr.com','{\"group_1\":\"/uploads/goods/1/1405444974.jpg\",\"group_2\":\"/uploads/goods/1/1405444977.jpg\",\"group_3\":\"/uploads/goods/1/1405444982.jpg\",\"large\":\"/uploads/goods/1/1405444985.jpg\",\"mini\":\"/uploads/goods/1/1405445006.jpg\"}','{\"cer_1\":\"/uploads/goods/1/1405445677.jpg\"}',1),(2,'义都农家米','米','<p>农家鲜米就是从农户那里收购最新鲜优质的稻谷，通过碾米机，将农家稻谷现场加工、现场销售，不做任何喷水、抛光、打蜡处理，也不添加任何漂白、防腐等化学添加剂，保持农家米的原汁原味，使大米的营养性、新鲜度和安全性都是最佳的，确保消费者粮食安全和营养。 河源义都农米磨坊采用农户直接进货，从源头上保证原料新鲜，确保了大米营养的最大化保留；整个过程直观清楚，营养健康看得见。古话说：&ldquo;五谷为养&rdquo;。粗细粮均有丰富的营养，搭配着吃才对健康更为有利。您的幸福安康是我们最大的心愿。金杯银杯不如您的口碑，真心希望义都农米的能成为你健康生活的伴侣。</p>\n','广东省梅州市五华县转水镇','1袋（15KG）',80.00,'广州稻草人','http://gzdcr.com','{}','{}',1),(3,'田园春通衢香粘米','米','','广东省河源市龙川县通衢镇','5KG',390.00,'一村一品','http://ycyp.com','{}','{}',1);
/*!40000 ALTER TABLE `tb_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_goods_category`
--

DROP TABLE IF EXISTS `tb_goods_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_goods_category` (
  `Name` varchar(128) NOT NULL,
  `Parent` varchar(128) NOT NULL,
  `Ordering` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods_category`
--

LOCK TABLES `tb_goods_category` WRITE;
/*!40000 ALTER TABLE `tb_goods_category` DISABLE KEYS */;
INSERT INTO `tb_goods_category` VALUES ('休闲食品','',254),('其它','',250),('冲调茶饮','',251),('油','粮油副食',0),('生鲜食品','',253),('米','粮油副食',0),('粮油副食','',255),('营养保健','',252);
/*!40000 ALTER TABLE `tb_goods_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_link`
--

DROP TABLE IF EXISTS `tb_link`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_link` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(128) NOT NULL,
  `Category` varchar(128) NOT NULL,
  `LinkUrl` varchar(256) NOT NULL,
  `ImageUrl` varchar(256) NOT NULL,
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_link`
--

LOCK TABLES `tb_link` WRITE;
/*!40000 ALTER TABLE `tb_link` DISABLE KEYS */;
INSERT INTO `tb_link` VALUES (1,'广州迅睿','友情链接','http://racent.com.cn','http://racent.com.cn',1),(2,'快雪时晴','友情链接','http://xianmau.me','http://xianmau.me',0);
/*!40000 ALTER TABLE `tb_link` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_link_category`
--

DROP TABLE IF EXISTS `tb_link_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_link_category` (
  `Name` varchar(128) NOT NULL,
  `Parent` varchar(128) NOT NULL,
  `Ordering` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_link_category`
--

LOCK TABLES `tb_link_category` WRITE;
/*!40000 ALTER TABLE `tb_link_category` DISABLE KEYS */;
INSERT INTO `tb_link_category` VALUES ('友情链接','',254),('首页图片轮播','',255);
/*!40000 ALTER TABLE `tb_link_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_origin`
--

DROP TABLE IF EXISTS `tb_origin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_origin` (
  `Name` varchar(128) NOT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_origin`
--

LOCK TABLES `tb_origin` WRITE;
/*!40000 ALTER TABLE `tb_origin` DISABLE KEYS */;
INSERT INTO `tb_origin` VALUES ('广东省梅州市五华县转水镇'),('广东省河源市龙川县义都镇'),('广东省河源市龙川县回龙镇'),('广东省河源市龙川县老隆镇'),('广东省河源市龙川县通衢镇'),('广东省韶关市乐昌县庆云镇');
/*!40000 ALTER TABLE `tb_origin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_role`
--

DROP TABLE IF EXISTS `tb_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_role` (
  `Name` varchar(128) NOT NULL,
  `Authority` text,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_role`
--

LOCK TABLES `tb_role` WRITE;
/*!40000 ALTER TABLE `tb_role` DISABLE KEYS */;
INSERT INTO `tb_role` VALUES ('root','[\"登录\",\"文章管理\",\"链接管理\",\"用户管理\",\"商店管理\",\"商品管理\",\"上传文件\",\"网站设置\",\"系统管理\",\"角色管理\",\"管理员管理\",\"货源管理\",\"帐号管理\"]'),('test','[\"登录\",\"角色管理\",\"管理员管理\",\"用户管理\",\"文章管理\",\"链接管理\",\"商店管理\",\"商品管理\",\"上传文件\",\"网站设置\",\"系统管理\",\"货源管理\"]'),('管理员','[\"登录\",\"角色管理\",\"管理员管理\",\"用户管理\",\"文章管理\",\"链接管理\",\"商店管理\",\"商品管理\",\"上传文件\",\"网站设置\",\"系统管理\",\"货源管理\"]');
/*!40000 ALTER TABLE `tb_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_shop`
--

DROP TABLE IF EXISTS `tb_shop`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_shop` (
  `Name` varchar(128) NOT NULL,
  `UserName` varchar(128) NOT NULL,
  `Kind` int(11) NOT NULL,
  `Introduce` text NOT NULL,
  `ApplyTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ApplyStatement` text NOT NULL,
  `Status` int(11) NOT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_shop`
--

LOCK TABLES `tb_shop` WRITE;
/*!40000 ALTER TABLE `tb_shop` DISABLE KEYS */;
INSERT INTO `tb_shop` VALUES ('绿稻人','gzdcr',0,'<p><img alt=\"\" src=\"/uploads/shop/绿稻人/1406637432.jpg\" style=\"height:905px; width:1133px\" /></p>\n','2014-07-29 12:37:39','<p>123</p>\n',0);
/*!40000 ALTER TABLE `tb_shop` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user` (
  `Name` varchar(128) NOT NULL,
  `Password` varchar(128) NOT NULL,
  `Level` int(11) NOT NULL DEFAULT '0',
  `BaseInfo` text NOT NULL,
  `ReceiveInfo` text NOT NULL,
  `RegisterTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `LoginHistory` text NOT NULL,
  `Status` int(11) DEFAULT '0',
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user`
--

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;
INSERT INTO `tb_user` VALUES ('xianmau','123456',2,'{\"Avatar\":\"http://xianmau.me\",\"RealName\":\"林仙茂\",\"Sex\":\"男\",\"Birthday\":\"1990-12-21\",\"Birthplace\":\"广东汕头\",\"Tel\":\"15913196613\",\"Email\":\"xianmaulin@gmail.com\"}','[{\"default\":\"false\",\"Address\":\"广东省广州市天河区\",\"Postcode\":\"510000\",\"Consignee\":\"xianmau\",\"Tel\":\"15913196613\",\"Remark\":\"Nothing\"},{\"default\":\"true\",\"Address\":\"广东省广州市天河区\",\"Postcode\":\"510000\",\"Consignee\":\"xianmao\",\"Tel\":\"15913196613\",\"Remark\":\"Nothing\"}]','2014-07-25 19:07:18','[\"2014-07-26 03:07:18\",\"2014-07-26 03:07:19\",\"2014-07-26 03:07:28\"]',2);
/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-08-04  0:51:56
