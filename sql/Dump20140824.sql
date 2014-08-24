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
INSERT INTO `tb_article` VALUES (1,'一村一品平台公测','新闻资讯','<p>123<img alt=\"\" src=\"/uploads/article/1408725853.jpg\" style=\"height:250px; width:250px\" /></p>\n','2014-07-20 15:41:27',0,1);
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
INSERT INTO `tb_article_category` VALUES ('健康食谱','',200),('商城公告','',254),('帮助中心','',253),('新闻资讯','',255),('服务承诺','',252),('系统内置','',0);
/*!40000 ALTER TABLE `tb_article_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_comment`
--

DROP TABLE IF EXISTS `tb_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_comment` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `GoodsId` int(11) NOT NULL,
  `Name` varchar(128) NOT NULL,
  `Content` text NOT NULL,
  `Time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_comment`
--

LOCK TABLES `tb_comment` WRITE;
/*!40000 ALTER TABLE `tb_comment` DISABLE KEYS */;
INSERT INTO `tb_comment` VALUES (1,23,'xianmaulin@gmail.com','很好很不错，下次还买，哈哈哈','2014-08-24 13:44:09',1),(2,23,'游客','啦啦啦我没有登录啦啦啦','2014-08-24 14:56:25',0);
/*!40000 ALTER TABLE `tb_comment` ENABLE KEYS */;
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
  `Recommend` text NOT NULL,
  `Content` text NOT NULL,
  `Origin` varchar(128) NOT NULL,
  `Unit` varchar(128) NOT NULL,
  `Price` decimal(8,2) NOT NULL DEFAULT '0.00',
  `Shop` varchar(128) NOT NULL,
  `BuyLink` varchar(256) NOT NULL,
  `Images` text NOT NULL,
  `Certificates` text NOT NULL,
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods`
--

LOCK TABLES `tb_goods` WRITE;
/*!40000 ALTER TABLE `tb_goods` DISABLE KEYS */;
INSERT INTO `tb_goods` VALUES (21,'新米东北大米香米','大米','','<p>// 请在此描述商品</p>\n','广东省梅州市五华县转水镇','500g',6.80,'一村一品','http://item.taobao.com/item.htm?spm=a1z10.1.w5866714-2795180462.4.RSL1GU&id=13518907798','{\"group_1\":\"/uploads/goods/21/1408718481.jpg\",\"group_2\":\"/uploads/goods/21/1408718484.jpg\",\"group_3\":\"/uploads/goods/21/1408718485.jpg\",\"group_4\":\"/uploads/goods/21/1408718488.jpg\",\"mini\":\"/uploads/goods/21/1408718491.jpg\",\"small\":\"/uploads/goods/21/1408718494.jpg\",\"medium\":\"/uploads/goods/21/1408718497.jpg\",\"large\":\"/uploads/goods/21/1408718500.jpg\"}','{}',0),(22,'九鲤湖 井冈山红米红稻米大米','大米','推荐理由写什么好呢，随便写长一点，但也不要太长，两百字左右吧，差不多四五行这样子。推荐理由写什么好呢，随便写长一点，但也不要太长，两百字左右吧，差不多四五行这样子。推荐理由写什么好呢，随便写长一点，但也不要太长，两百字左右吧，差不多四五行这样子。推荐理由写什么好呢，随便写长一点，但也不要太长，两百字左右吧，差不多四五行这样子。','<p>// 请在此描述商品</p>\n','广东省河源市龙川县义都镇','500g',7.80,'一村一品','http://detail.tmall.com/item.htm?spm=a1z10.3.w4011-4136801490.85.ams8WS&id=35393983208&rn=7d01d43210b7ea99bc677433e2fe9c00','{\"group_1\":\"/uploads/goods/22/1408719416.jpg\",\"group_2\":\"/uploads/goods/22/1408719419.jpg\",\"group_3\":\"/uploads/goods/22/1408719422.jpg\",\"group_4\":\"/uploads/goods/22/1408719426.jpg\",\"mini\":\"/uploads/goods/22/1408719433.jpg\",\"small\":\"/uploads/goods/22/1408719436.jpg\",\"medium\":\"/uploads/goods/22/1408719440.jpg\",\"large\":\"/uploads/goods/22/1408719443.jpg\"}','{}',0),(23,'北大荒长粒香米5公斤装（演示商品）','大米','北大荒长粒香米，米粒修长，皎莹如玉，气味醇正芳香。其产地黑龙江省大庆市肇源农场，地理位置优越，气候怡人，日照长，大米比普通米含有更多的糖类、蛋白质、膳食纤维叶酸等生物活性物质。其香味浓郁，饭粒油亮，出饭率高，粘性较小，米质较脆，是一款难得的好米，一村一品极力推荐。','<p style=\"text-align:center\"><img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_1.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_2.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_3.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_4.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_5.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_6.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_7.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_8.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_9.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_10.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_11.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_12.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_13.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_14.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_15.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_16.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_17.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_18.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_19.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_20.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_21.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_22.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_23.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_24.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_25.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_26.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_27.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_28.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_29.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_30.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_31.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_32.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_33.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_34.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_35.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_36.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_37.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_38.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_39.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_40.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_41.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_42.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_43.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_44.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_45.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_46.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_47.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_48.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_49.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_50.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_51.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_52.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_53.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_54.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_55.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_56.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_57.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_58.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_59.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_60.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_61.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_62.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_63.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_64.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_65.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_66.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_67.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_68.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_69.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /><br />\n<img alt=\"\" class=\"lazy\" src=\"http://p01.sfbest.com/1300049822/1406537801_70.jpg\" style=\"border:0px; color:rgb(102, 102, 102); display:inline; font-family:arial,verdana,宋体; font-size:14px; line-height:24px; margin:0px; padding:0px; text-align:center\" /></p>\n','广东省韶关市乐昌县庆云镇','5.00kg',56.00,'一村一品','http://www.sfbest.com/html/products/50/1300049822.html?etc_n=padnet&etc_s=yqf&union_source=CPS&union_medium=yqf&union_campaign=cps&utm_source=cps','{\"group_1\":\"/uploads/goods/23/1408876827.jpg\",\"group_2\":\"/uploads/goods/23/1408876836.jpg\",\"group_3\":\"/uploads/goods/23/1408876840.jpg\",\"group_4\":\"/uploads/goods/23/1408876846.jpg\",\"medium\":\"/uploads/goods/23/1408876973.jpg\"}','{}',1);
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
INSERT INTO `tb_goods_category` VALUES ('优质粮油','',255),('其它','',0),('农家干货','',254),('华农出品','',252),('大米','优质粮油',239),('干菜','农家干货',228),('新鲜果蔬','',253),('木耳','农家干货',230),('杂粮','优质粮油',236),('枣子','农家干货',226),('水果','新鲜果蔬',220),('茶叶','华农出品',207),('蔬菜','新鲜果蔬',209),('酸奶','华农出品',210),('银耳','农家干货',229),('面条','优质粮油',237),('面粉','优质粮油',238),('食用油','优质粮油',240),('香菇','农家干货',227),('鸡蛋','华农出品',208);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
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

-- Dump completed on 2014-08-24 22:59:42
