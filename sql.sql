# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 8.0.22)
# Database: go_gin
# Generation Time: 2022-03-03 06:15:03 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table articles
# ------------------------------------------------------------

CREATE TABLE `articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;

INSERT INTO `articles` (`id`, `name`, `user_id`, `title`, `content`, `date`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'名称',NULL,'标题','内容','2022-01-17 12:12:12','2022-01-17 12:12:12','2022-01-17 12:12:12',NULL),
	(2,'名称',0,'标题','内容1',NULL,'2022-01-18 11:29:54','2022-01-18 13:38:28',NULL),
	(3,'名称',0,'标题','内容',NULL,'2022-01-18 13:40:07','2022-01-18 13:40:07',NULL),
	(4,'名称',0,'标题','内容1',NULL,'2022-01-18 13:40:10','2022-01-18 13:47:45',NULL);

/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table articles_category
# ------------------------------------------------------------

CREATE TABLE `articles_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `level` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `articles_category` WRITE;
/*!40000 ALTER TABLE `articles_category` DISABLE KEYS */;

INSERT INTO `articles_category` (`id`, `name`, `level`, `parent_id`, `created_at`, `updated_at`)
VALUES
	(1,'文章一级类目1',1,0,'2022-02-18 17:26:02','2022-02-18 17:52:51'),
	(2,'文章二级类目1',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26'),
	(3,'文章二级类目2',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26'),
	(4,'文章二级类目3',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26'),
	(5,'文章一级类目2',1,0,'2022-02-18 17:26:02','2022-02-18 17:52:51'),
	(6,'文章一级类目0221',1,0,'2022-02-21 16:05:35','2022-02-21 16:05:35'),
	(7,'文章二级类目0221-1',2,6,'2022-02-21 16:06:08','2022-02-21 16:06:08'),
	(8,'一级类目-test',1,0,'2022-02-21 16:12:26','2022-02-21 16:12:26'),
	(9,'二级类目-test-1',2,8,'2022-02-21 16:12:26','2022-02-21 16:12:26'),
	(10,'二级类目-test-2',2,8,'2022-02-21 16:12:26','2022-02-21 16:12:26'),
	(11,'二级类目-test-3',2,8,'2022-02-21 16:12:26','2022-02-21 16:12:26');

/*!40000 ALTER TABLE `articles_category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table articles_category_bak
# ------------------------------------------------------------

CREATE TABLE `articles_category_bak` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `level` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `articles_category_bak` WRITE;
/*!40000 ALTER TABLE `articles_category_bak` DISABLE KEYS */;

INSERT INTO `articles_category_bak` (`id`, `name`, `level`, `parent_id`, `created_at`, `updated_at`)
VALUES
	(1,'文章一级类目1',1,0,'2022-02-18 17:26:02','2022-02-18 17:52:51'),
	(2,'文章二级类目1',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26'),
	(3,'文章二级类目2',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26'),
	(4,'文章二级类目3',2,1,'2022-02-18 17:26:26','2022-02-18 17:26:26');

/*!40000 ALTER TABLE `articles_category_bak` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table articles_comments
# ------------------------------------------------------------

CREATE TABLE `articles_comments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `comments` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# Dump of table articles_logs
# ------------------------------------------------------------

CREATE TABLE `articles_logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `articles_id` int DEFAULT NULL,
  `manager_id` int DEFAULT NULL,
  `logs` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# Dump of table managers
# ------------------------------------------------------------

CREATE TABLE `managers` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `managers` WRITE;
/*!40000 ALTER TABLE `managers` DISABLE KEYS */;

INSERT INTO `managers` (`id`, `name`, `phone`, `pwd`, `email`, `created_at`, `updated_at`)
VALUES
	(1,'admin','123456798','25f9e794323b453885f5181f1b624d0b','email@123.com','2022-01-21 12:09:50','2022-01-21 12:09:50');

/*!40000 ALTER TABLE `managers` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table test0301
# ------------------------------------------------------------

CREATE TABLE `test0301` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `test0301` WRITE;
/*!40000 ALTER TABLE `test0301` DISABLE KEYS */;

INSERT INTO `test0301` (`id`, `name`, `time`)
VALUES
	(1,NULL,'2022-03-01 18:43:31');

/*!40000 ALTER TABLE `test0301` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `sex` int DEFAULT NULL,
  `articles_count` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
