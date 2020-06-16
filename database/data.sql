-- MySQL dump 10.13  Distrib 8.0.13, for osx10.14 (x86_64)
--
-- Host: localhost    Database: test2
-- ------------------------------------------------------
-- Server version	8.0.13

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `games`
--

DROP TABLE IF EXISTS `games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `games` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `cover` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `release_date` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `rating` decimal(10,2) DEFAULT NULL,
  `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `languages` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `platforms` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `douban_id` int(11) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT '0.00',
  `quantity` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_games_title` (`title`),
  KEY `idx_games_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `games`
--

LOCK TABLES `games` WRITE;
/*!40000 ALTER TABLE `games` DISABLE KEYS */;
INSERT INTO `games` VALUES (1,'2020-06-22 04:26:47','2020-06-23 00:50:06',NULL,'马里奥派对','pd.jpg','2018-10-05',9.10,'美版','简体中文','Switch',30245974,309.00,4),(2,'2020-06-22 04:30:05','2020-06-23 00:50:25',NULL,'马里奥奥德赛','ads.jpg','2017-10-27',9.50,'澳版','简体中文','Switch',26954652,305.00,5),(3,'2020-06-22 04:31:15','2020-06-23 00:50:39',NULL,'马里奥创作家2','czj2.jpg','2019-06-28',9.40,'美版','简体中文','Switch',30459464,305.00,5),(4,'2020-06-22 04:50:28','2020-06-23 00:50:54',NULL,'我的世界','wdsj.jpg','2011-11-18',9.30,'美版','简体中文','Switch',10734157,179.00,1),(5,'2020-06-22 04:51:32','2020-06-23 00:51:11',NULL,'集合吧动物森友会','ds.jpg','2020-03-20',9.60,'澳版','简体中文','Switch',30325263,330.00,10),(6,'2020-06-22 04:52:29','2020-06-23 00:51:29',NULL,'舞力全开2020','wlqk.jpg','2019-11-05',8.70,'美版','简体中文','Switch',34429191,209.00,10),(7,'2020-06-22 04:54:35','2020-06-23 00:51:51',NULL,'文明6','wm6.jpg','2016-10-21',9.20,'美版','简体中文','Switch',26791492,130.00,5),(8,'2020-06-22 04:55:49','2020-06-23 00:44:14',NULL,'上古卷轴5:天际','sgjz.jpg','2011-11-11',9.40,'澳版','简体中文','Switch',10735290,249.00,10),(9,'2020-06-22 04:57:10','2020-06-23 00:52:06',NULL,'暗黑破坏神3','ahphs3.jpg','2012-05-15',8.20,'澳版','简体中文','Switch',10729903,275.00,2),(10,'2020-06-22 04:58:20','2020-06-23 00:52:22',NULL,'动森限定保护包','bao.jpg','2020-03-20',10.00,'日版','无中文','Switch',0,180.00,4),(11,'2020-06-22 04:59:03','2020-06-23 00:52:33',NULL,'动森限定Switch主机','switch.jpg','2020-03-20',10.00,'美版','简体中文','Switch',0,3300.00,1),(12,'2020-06-22 05:00:28','2020-06-23 00:52:47',NULL,'健身环大冒险','jsh.jpg','2019-10-18',9.50,'美版/欧版/澳版','简体中文','Switch',34824349,749.00,1),(13,'2020-06-22 05:04:20','2020-06-23 00:53:01',NULL,'异度神剑终极版欧版典藏套装','ydsj.jpg','2020-05-29',10.00,'欧版','简体中文','Switch',0,1299.00,1),(14,'2020-06-22 05:09:33','2020-06-23 00:53:26',NULL,'赛博朋克2077典藏套装','cyberpunk2077.jpg','2020-11-19',10.00,'美版/欧版','无中文','PS4',0,2900.00,4);
/*!40000 ALTER TABLE `games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subscriptions`
--

DROP TABLE IF EXISTS `subscriptions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
 CREATE TABLE `subscriptions` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `created_at` datetime DEFAULT NULL,
 `updated_at` datetime DEFAULT NULL,
 `deleted_at` datetime DEFAULT NULL,
 `uid` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
 `gid` int(11) DEFAULT NULL,
 PRIMARY KEY (`id`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subscriptions`
--

LOCK TABLES `subscriptions` WRITE;
/*!40000 ALTER TABLE `subscriptions` DISABLE KEYS */;
/*!40000 ALTER TABLE `subscriptions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-24 12:58:35
