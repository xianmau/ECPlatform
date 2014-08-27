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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_article`
--

LOCK TABLES `tb_article` WRITE;
/*!40000 ALTER TABLE `tb_article` DISABLE KEYS */;
INSERT INTO `tb_article` VALUES (1,'一村一品平台公测','新闻资讯','<p>123<img alt=\"\" src=\"/uploads/article/1408725853.jpg\" style=\"height:250px; width:250px\" /></p>\n','2014-07-20 15:41:27',0,1),(2,'科学使用健康食谱，快速摆脱节日综合症','健康食谱','<p>春节长假刚刚结束，各种年后综合征也随之而来。三元营养专家日前提醒大家，春节后首先要调养肠胃，通过饮食让身体尽快恢复状态。同时，要科学补充多种身体必需物质，提高脑力和体力劳动效率，以良好的状态从容应对接下来忙碌而紧张的工作。</p>\n\n<p>善用健康食谱，快速调整身体状态</p>\n\n<p>三元营养专家建议，节后几天要减少米、面、糖果、甜糕点的摄入量，主食最好以粗粮为主，如玉米、燕麦、小米、豆类等，同时吃些深色或绿色的蔬菜，并多喝粥和汤。蔬菜最好生食或做成蔬菜粥和汤，这样不仅能补充身体所需的纤维素，而且利于排毒，让紊乱的胃肠得到休息、调整。</p>\n\n<p>有选择地多吃水果也能有效调整失衡的消化功能。如饭后1小时喝橙汁和木瓜汁，可调整消化机能；吃苹果和香蕉可促进消化。在年后几天还要强迫自己多喝白开水或绿茶。这样可以加快胃肠道的新陈代谢，减轻大量肉类食物和酒精对肝脏的危害。</p>\n\n<p>三元营养专家同时提醒大家，不要为了尽快平衡油腻而一味的食&ldquo;素&rdquo;，那样对胃肠道同样不利，不仅难以达到&ldquo;休整&rdquo;目的，甚至还会&ldquo;雪上加霜&rdquo;。</p>\n\n<p>如燕麦等食物富含膳食纤维，并且含有相当数量的不可溶性粗纤维，一方面会延缓胃排空，导致腹胀、早饱、消化不良、食欲降低等症状，甚至还可能影响下一餐的进食。</p>\n\n<p>其实，从节日的大鱼大肉一下改为以素为主，会造成新一轮的营养失衡，特别是动物食品里富含的B族维生素、微量元素以及优质蛋白、必需氨基酸等，会因此摄入不足。</p>\n\n<p>平衡营养和膳食，提高脑力劳动效率</p>\n\n<p>因此，针对节日饮食使营养失衡的情况，三元营养专家特地表示，应注意营养和膳食平衡，使自己有充沛的精力和健康的身体以胜任各项工作。</p>\n\n<p>注意减少脂肪摄入。一般情况下，应减少脂肪摄入量，少吃油炸食品，以防超重和肥胖。脂肪的摄入量标准应占总热能的20%～25%，如果脂肪摄入过多，容易导致活动能力降低，影响工作效率。</p>\n\n<p>维生素摄入要充足。维生素是维持生理功能的重要因素，特别是与脑和神经代谢有关的维生素，如维生素B1、维生素B6等。另外，抗氧化营养素如&beta;胡萝卜素、维生素C、维生素E等，都有利于提高工作效率，各种新鲜蔬菜和水果中其含量尤为丰富。白领人士工作繁忙，饮食中的维生素营养常被忽略。因此，不妨常补充一些奶制品，来保证维生素的均衡水平。还应多摄入一些钙、镁、锌和铁，以提高脑力劳动的效率。</p>\n\n<p>注意补充氨基酸。白领人士的工作特点是用脑，因此营养脑神经的氨基酸供给要充足。脑组织中的游离氨基酸含量以谷氨酸为最高，其次是天门冬氨酸等。豆类、芝麻等含谷氨酸及天门冬氨酸较丰富，应适当多吃。</p>\n\n<p>科学食用奶制品，补充营养更全面</p>\n\n<p>三元营养专家建议白领人士多食用奶制品。牛奶味甘、性微寒，牛奶中丰富的蛋白质能有效缓解身体的疲劳感；牛奶中的乳糖还可以促进人体对钙和铁的吸收，增强肠蠕动，润肠通便；牛奶还可以阻止人体吸收食物中有毒的金属铅和镉，其生物活性物质SOD能清除体内有害物质，增强免疫力，减少肝脏的负担。</p>\n\n<p>不仅是牛奶，奶粉也同样能提供人们身体需要的多种营养物质。据了解，三元的成人奶粉特别打造了针对中国人体质需求的ilactou营养系统，其益智营养因子中富含的乳铁蛋白可促进铁吸收；其促进吸收营养因子除合理调整钙磷比例外，还富含维生素D3、镁。</p>\n\n<p><br />\n三元的奶粉中还特别添加益生元组合，防止上火的同时，还能促进肠道内有益菌群增殖，抑制有害菌生长，清楚肠内的有害物质，保持肠内的菌群平衡，促进营养吸收，达到改善健康的目的。还含有丰富的维生素A、维生素C、维生素E和B族维生素，钙、铜、铁、锌等微量元素。钙可以抑制神经的紧张亢奋，使人心情舒缓；铁、铜和卵磷脂能大大提高大脑的工作效率；优质蛋白质还能有效提高人体的免疫系统功能。</p>\n\n<p>三元营养专家表示，奶制品无论是牛奶还是奶粉，现在都被营养学界看成是对抗&ldquo;节后综合症&rdquo;的&ldquo;良药&rdquo;。因此，人们在日常饮食中，要多注意食用奶制品，全面提高身体素质。</p>\n\n<p>(本文来源：网易 )&nbsp;</p>\n','2014-08-25 17:43:11',0,1),(3,'“乐活族”：借助营养标签打造健康食谱','健康食谱','<p>&nbsp; &nbsp;&nbsp;近日，记者发现，一种凸显人性化的图形化营养标签已悄然上市。该标签较以往表格形式更加清晰、易懂，为崇尚均衡膳食的&ldquo;乐活族&rdquo;提供了方便的&ldquo;营养指南&rdquo;，一出现便备受热捧。</p>\n\n<p>&nbsp;&nbsp;&nbsp; &ldquo;乐活族&rdquo;们信奉健康、自然、可持续的生活方式。 &ldquo;经常运动、适度休息、均衡饮食，不把健康的责任丢给医生&rdquo;，是乐活族十大宣言之一。所谓平衡饮食，是指饮食中各种营养素，包括蛋白质、脂肪、糖类、无机盐和维生素等，种类齐全，数量充足，比例适当。任何一种过多或过少都会给健康带来危害。在追求健康、崇尚自然的道路上，学习合理搭配膳食、均衡营养是实践&ldquo;乐活&rdquo;的必修课。秋冬季节正是美食的季节，&ldquo;乐活族&rdquo;们经常为如何均衡膳食大伤脑筋，现在，他们中许多人已经找到了自己的专属&ldquo;营养顾问&rdquo;&mdash; 食品营养标签。营养标签，即预包装食品的包装上标注的产品营养信息。根据卫生部的《食品营养标签管理规范》，预包装食品的包装上应按规定标明产品中能量以及蛋白质、脂肪、碳水化合物、钠等四种核心营养素的含量以及营养素占&ldquo;营养素参考值&rdquo;的百分比（NRV%）。如果说营养标签是&ldquo;乐活族&rdquo;的营养顾问，那么NRV%就是这位顾问手中的一把标尺，帮助人们了解食品中主要营养成分占每日建议摄入量的比例。有了营养标签，&ldquo;乐活族&rdquo;们就能轻松地控制营养素的摄入，实现膳食平衡。</p>\n\n<p><br />\n&nbsp;&nbsp;&nbsp; 以德芙巧克力为例，最新上市的德芙巧克力营养标签全面升级，除国家规定的营养成分表外，还使用了图形化的营养标签；该标签不仅按照每100克食品所含的营养成分进行了说明，同时按照更实用的单位如每包装标注了营养信息，消费者据此可以得知所享用的巧克力究竟能提供多少营养和能量。</p>\n\n<p>&nbsp;&nbsp;&nbsp; 如图所示，吃一块43克德芙巧克力可以获取一天所需的12%的能量。这无疑对消费者选购食物具有很好的参考价值，特别是对崇尚饮食均衡的乐活族。运用好营养标签，我们大家都可以像&ldquo;乐活族&rdquo;一样，在多吃蔬菜水果的同时也不会拒绝用香浓丝滑的巧克力犒劳自己，补充每日所需的营养和能量。玛氏集团这种将消费者的健康置于首位的精神，与&ldquo;乐活族&rdquo;崇尚的生活理念不谋而合，同时图形化营养标签也将成为消费者均衡饮食和健康生活的好帮手。</p>\n\n<p>(本文来源：网易 )&nbsp;</p>\n','2014-08-25 18:07:40',0,1);
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
-- Table structure for table `tb_goods_category_info`
--

DROP TABLE IF EXISTS `tb_goods_category_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_goods_category_info` (
  `CategoryName` varchar(128) NOT NULL,
  `CategoryInfo` text NOT NULL,
  `CategoryImage` text NOT NULL,
  PRIMARY KEY (`CategoryName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_goods_category_info`
--

LOCK TABLES `tb_goods_category_info` WRITE;
/*!40000 ALTER TABLE `tb_goods_category_info` DISABLE KEYS */;
INSERT INTO `tb_goods_category_info` VALUES ('优质粮油','/uploads/other/1.png','/uploads/other/2.png');
/*!40000 ALTER TABLE `tb_goods_category_info` ENABLE KEYS */;
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
-- Table structure for table `tb_message`
--

DROP TABLE IF EXISTS `tb_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_message` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(128) NOT NULL,
  `Email` varchar(128) NOT NULL,
  `Content` text NOT NULL,
  `Time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_message`
--

LOCK TABLES `tb_message` WRITE;
/*!40000 ALTER TABLE `tb_message` DISABLE KEYS */;
INSERT INTO `tb_message` VALUES (1,'xianmau','xianmaulin@gmail.com','good','2014-08-25 16:54:16',0);
/*!40000 ALTER TABLE `tb_message` ENABLE KEYS */;
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
INSERT INTO `tb_shop` VALUES ('绿稻人','gzdcr',0,'<p>绿稻人</p>\n','2014-08-24 15:06:36','<p>无证</p>\n',1);
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

-- Dump completed on 2014-08-28  0:46:03
