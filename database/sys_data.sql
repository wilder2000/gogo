-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: yhdb
-- ------------------------------------------------------
-- Server version	5.7.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `s_urlmappings`
--

LOCK TABLES `s_urlmappings` WRITE;
/*!40000 ALTER TABLE `s_urlmappings` DISABLE KEYS */;
INSERT INTO `s_urlmappings` VALUES (1,10,'/pwd'),(2,10,'/cpwd'),(3,10,'/uquery'),(4,10,'/rquery'),(5,10,'/radd'),(6,10,'/ug'),(7,10,'/rm'),(8,10,'/dm'),(9,10,'/um'),(10,10,'/upro'),(11,10,'/mif/c'),(12,10,'/mif/q'),(13,10,'/mif/d'),(14,10,'/mif/u'),(15,10,'/avatorup'),(16,11,'/avatorup');
/*!40000 ALTER TABLE `s_urlmappings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_role`
--

LOCK TABLES `s_role` WRITE;
/*!40000 ALTER TABLE `s_role` DISABLE KEYS */;
INSERT INTO `s_role` VALUES (1,'终端用户角色','2024-02-28 22:12:21'),(2,'管理员角色','2024-03-01 20:31:48');
/*!40000 ALTER TABLE `s_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_operators`
--

LOCK TABLES `s_operators` WRITE;
/*!40000 ALTER TABLE `s_operators` DISABLE KEYS */;
INSERT INTO `s_operators` VALUES (10,'管理功能'),(11,'查看功能');
/*!40000 ALTER TABLE `s_operators` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_items`
--

LOCK TABLES `s_items` WRITE;
/*!40000 ALTER TABLE `s_items` DISABLE KEYS */;
INSERT INTO `s_items` VALUES (1,'冰糕就蒜',1),(2,'巡山小妖精',1),(3,'再见亦不见',1),(4,'难以抑制的思念',1),(5,'凌芹',1),(6,'我想要の快乐并不多',1),(7,'叽喱咕噜︶',1),(8,'▼心奴',1),(9,'有沒有人敢陪我到老',1),(10,'胸有大痣',1),(11,'路尽隐香处',1),(12,'莫语',1),(13,'不贱不开心',1),(14,'鱼塘空荡荡海王在冲浪',1),(15,'冷心冷面',1),(16,'だ简ゑ箪ā爱',1),(17,'丶浅瞳°',1),(18,'爱。归零',1),(19,'流星划过黑暗的夜空つ',1),(20,'中国移动，哥不动',1),(21,'乜许会寂寞',1),(22,'无鞋的呆呆',1),(23,'娇嗔语气',1),(24,'小乔躲猫猫',1),(25,'你的快乐已到期请及时续费',1),(26,'不风流怎样倜傥',1);
/*!40000 ALTER TABLE `s_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_resources`
--

LOCK TABLES `s_resources` WRITE;
/*!40000 ALTER TABLE `s_resources` DISABLE KEYS */;
/*!40000 ALTER TABLE `s_resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_rolegroup`
--

LOCK TABLES `s_rolegroup` WRITE;
/*!40000 ALTER TABLE `s_rolegroup` DISABLE KEYS */;
INSERT INTO `s_rolegroup` VALUES (1,1),(1,2),(2,2);
/*!40000 ALTER TABLE `s_rolegroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_departments`
--

LOCK TABLES `s_departments` WRITE;
/*!40000 ALTER TABLE `s_departments` DISABLE KEYS */;
INSERT INTO `s_departments` VALUES (1,'运维部门','2024-02-28 22:13:06',NULL);
/*!40000 ALTER TABLE `s_departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_groupuser`
--

LOCK TABLES `s_groupuser` WRITE;
/*!40000 ALTER TABLE `s_groupuser` DISABLE KEYS */;
INSERT INTO `s_groupuser` VALUES (1,'A74B460E-6077-4990-BD7E-B1A071D6CDD3'),(2,'urn:uuid:68c46ecc-d4b6-11ee-9ed1-525400e447da');
/*!40000 ALTER TABLE `s_groupuser` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_roleoperator`
--

LOCK TABLES `s_roleoperator` WRITE;
/*!40000 ALTER TABLE `s_roleoperator` DISABLE KEYS */;
INSERT INTO `s_roleoperator` VALUES (1,11,0),(2,10,0),(2,11,0);
/*!40000 ALTER TABLE `s_roleoperator` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_group`
--

LOCK TABLES `s_group` WRITE;
/*!40000 ALTER TABLE `s_group` DISABLE KEYS */;
INSERT INTO `s_group` VALUES (1,'终端用户组','2024-03-01 13:30:25'),(2,'管理员组','2024-03-01 20:32:11');
/*!40000 ALTER TABLE `s_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_depusers`
--

LOCK TABLES `s_depusers` WRITE;
/*!40000 ALTER TABLE `s_depusers` DISABLE KEYS */;
/*!40000 ALTER TABLE `s_depusers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_users`
--

LOCK TABLES `s_users` WRITE;
/*!40000 ALTER TABLE `s_users` DISABLE KEYS */;
INSERT INTO `s_users` VALUES ('urn:uuid:68c46ecc-d4b6-11ee-9ed1-525400e447da','wild.shang@163.com','$2a$04$BRAGWesq3kmcQ/KPatE47OEhRWAsrfAf/lCljwuGn5me7oaUGlnWy','流星划过黑暗的夜空つ','',0,'2024-02-26 22:50:40','2024-02-26 22:50:40',2,'','管理员');
/*!40000 ALTER TABLE `s_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-03-01 22:24:06
