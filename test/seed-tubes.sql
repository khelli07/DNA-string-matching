-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: gosql
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `diagnoses`
--

DROP TABLE IF EXISTS `diagnoses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `diagnoses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `input_date` longtext DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `dna_sequence` longtext DEFAULT NULL,
  `disease_name` longtext DEFAULT NULL,
  `percentage` float DEFAULT NULL,
  `result` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `diagnoses`
--

LOCK TABLES `diagnoses` WRITE;
/*!40000 ALTER TABLE `diagnoses` DISABLE KEYS */;
INSERT INTO `diagnoses` VALUES (1,'2022-04-19','Orang','CATGATCGATGCCATGATCATGATCGATGC','COV',0.8,1),(3,'2022-04-23','Orang','CAGTAC','Polio',0.8,1),(4,'2022-04-28','nama orang','AGCACTGCAAGCTAGCATCGAT','test1',100,1),(5,'2022-04-28','nama orang 2','AGTACTGCAAGCTAGCATCGAT','test1',100,1),(6,'2022-04-28','nama orang 2','CGCCGTCGCTCGTGCTCGTGCTC','test1',66,0),(7,'2022-04-28','nama orang 2','CGCCGTCGCTCGTGCTCGTGCTC','test1',66,0),(8,'2022-04-28','nama orang 2','CGCCGTCGCTCGTGCTCGTGCTC','test1',66,0),(9,'2022-04-28','nama orang 2','CGCCGTCGCTCGTGCTCGTGCTC','test1',66,0),(10,'2022-04-28','nama orang 2','CGCCGTCGCTCGTGCTCGTGCTC','test1',66,0);
/*!40000 ALTER TABLE `diagnoses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `diseases`
--

DROP TABLE IF EXISTS `diseases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `diseases` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(191) DEFAULT NULL,
  `pattern` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `name_2` (`name`),
  UNIQUE KEY `name_3` (`name`),
  UNIQUE KEY `name_4` (`name`),
  UNIQUE KEY `name_5` (`name`),
  UNIQUE KEY `name_6` (`name`),
  UNIQUE KEY `name_7` (`name`),
  UNIQUE KEY `name_8` (`name`),
  UNIQUE KEY `name_9` (`name`),
  UNIQUE KEY `name_10` (`name`),
  UNIQUE KEY `name_11` (`name`),
  UNIQUE KEY `name_12` (`name`),
  UNIQUE KEY `name_13` (`name`),
  UNIQUE KEY `name_14` (`name`),
  UNIQUE KEY `name_15` (`name`),
  UNIQUE KEY `name_16` (`name`),
  UNIQUE KEY `name_17` (`name`),
  UNIQUE KEY `name_18` (`name`),
  UNIQUE KEY `name_19` (`name`),
  UNIQUE KEY `name_20` (`name`),
  UNIQUE KEY `name_21` (`name`),
  UNIQUE KEY `name_22` (`name`),
  UNIQUE KEY `name_23` (`name`),
  UNIQUE KEY `name_24` (`name`),
  UNIQUE KEY `name_25` (`name`),
  UNIQUE KEY `name_26` (`name`),
  UNIQUE KEY `name_27` (`name`),
  UNIQUE KEY `name_28` (`name`),
  UNIQUE KEY `name_29` (`name`),
  UNIQUE KEY `name_30` (`name`),
  UNIQUE KEY `name_31` (`name`),
  UNIQUE KEY `name_32` (`name`),
  UNIQUE KEY `name_33` (`name`),
  UNIQUE KEY `name_34` (`name`),
  UNIQUE KEY `name_35` (`name`),
  UNIQUE KEY `name_36` (`name`),
  UNIQUE KEY `name_37` (`name`),
  UNIQUE KEY `name_38` (`name`),
  UNIQUE KEY `name_39` (`name`),
  UNIQUE KEY `name_40` (`name`),
  UNIQUE KEY `name_41` (`name`),
  UNIQUE KEY `name_42` (`name`),
  UNIQUE KEY `name_43` (`name`),
  UNIQUE KEY `name_44` (`name`),
  UNIQUE KEY `name_45` (`name`),
  UNIQUE KEY `name_46` (`name`),
  UNIQUE KEY `name_47` (`name`),
  UNIQUE KEY `name_48` (`name`),
  UNIQUE KEY `name_49` (`name`),
  UNIQUE KEY `name_50` (`name`),
  UNIQUE KEY `name_51` (`name`),
  UNIQUE KEY `name_52` (`name`),
  UNIQUE KEY `name_53` (`name`),
  UNIQUE KEY `name_54` (`name`),
  UNIQUE KEY `name_55` (`name`),
  UNIQUE KEY `name_56` (`name`),
  UNIQUE KEY `name_57` (`name`),
  UNIQUE KEY `name_58` (`name`),
  UNIQUE KEY `name_59` (`name`),
  UNIQUE KEY `name_60` (`name`),
  UNIQUE KEY `name_61` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `diseases`
--

LOCK TABLES `diseases` WRITE;
/*!40000 ALTER TABLE `diseases` DISABLE KEYS */;
INSERT INTO `diseases` VALUES (1,'PENYAKIT-1','CATGATCGATGC'),(3,'COV','ACCTGAAGCTA'),(8,'Polio','CATGATGCATGCATACCTG'),(9,'test1','AGC');
/*!40000 ALTER TABLE `diseases` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-29 23:22:52
