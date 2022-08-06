mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- Connecting to db...
-- MySQL dump 10.13  Distrib 8.0.30, for Linux (x86_64)
--
-- Host: db    Database: dclib_test
-- ------------------------------------------------------
-- Server version	8.0.30

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
-- Retrieving table structure for table authors...

--
-- Table structure for table `authors`
--

DROP TABLE IF EXISTS `authors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authors` (
  `author_id` int NOT NULL AUTO_INCREMENT,
  `author_name` varchar(100) NOT NULL,
  `author_surname` varchar(100) NOT NULL,
  `author_patrynomic` varchar(100) DEFAULT NULL,
  `author_photo` varchar(100) NOT NULL,
  `author_stars` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `authors`
--

-- Retrieving rows...
LOCK TABLES `authors` WRITE;
/*!40000 ALTER TABLE `authors` DISABLE KEYS */;
INSERT INTO `authors` VALUES (1,'Lev','Tolstoy','Nickolaevich','',0),(2,'Steven','King','Edwid','',0),(3,'Adolph','Hitler','Alaizovich','',0),(4,'Mikhail','Lermontov','Jurievich','',0),(5,'Nikolay','Hohol','Vasilevich','',0),(6,'Fedor','Dostoevsky','Mikhailovich','',0),(7,'Ray','Bradbury','Douglas','',0),(8,'Vladimir','Lenin','Ilich','',0);
/*!40000 ALTER TABLE `authors` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table booking...

--
-- Table structure for table `booking`
--

DROP TABLE IF EXISTS `booking`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `booking` (
  `id` int NOT NULL AUTO_INCREMENT,
  `book_id` int NOT NULL,
  `userid` int NOT NULL,
  `date_of_issue` timestamp NOT NULL,
  `date_of_delivery` timestamp NULL DEFAULT NULL,
  `is_confirm` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `book_id` (`book_id`),
  KEY `userid` (`userid`),
  CONSTRAINT `booking_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`),
  CONSTRAINT `booking_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `booking`
--

-- Retrieving rows...
LOCK TABLES `booking` WRITE;
/*!40000 ALTER TABLE `booking` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table books...

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `book_id` int NOT NULL AUTO_INCREMENT,
  `book_name` varchar(100) NOT NULL,
  `book_count` int DEFAULT NULL,
  `book_photo` varchar(100) NOT NULL,
  `book_stars` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `books`
--

-- Retrieving rows...
LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'War and peace',1,'',0),(2,'It',2,'',0),(3,'Something',1,'',0),(4,'Mein Kamph',1488,'',0),(5,'Hero of our time',1,'',0),(6,'Taras Bulba',2,'',0),(7,'Died souls',3,'',0),(8,'Crime and punishment',1,'',0),(9,'Dandelion wine',2,'',0),(10,'The State and the Revolution',5,'',0);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table books_authors...

--
-- Table structure for table `books_authors`
--

DROP TABLE IF EXISTS `books_authors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books_authors` (
  `book_id` int NOT NULL,
  `author_id` int NOT NULL,
  KEY `book_id` (`book_id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `books_authors_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`),
  CONSTRAINT `books_authors_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `authors` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `books_authors`
--

-- Retrieving rows...
LOCK TABLES `books_authors` WRITE;
/*!40000 ALTER TABLE `books_authors` DISABLE KEYS */;
INSERT INTO `books_authors` VALUES (1,1),(2,2),(3,2),(4,3),(5,4),(6,5),(7,5),(8,6),(9,7),(10,8);
/*!40000 ALTER TABLE `books_authors` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table favoriete_authors...

--
-- Table structure for table `favoriete_authors`
--

DROP TABLE IF EXISTS `favoriete_authors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `favoriete_authors` (
  `userid` int NOT NULL,
  `author_id` int NOT NULL,
  KEY `userid` (`userid`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `favoriete_authors_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`),
  CONSTRAINT `favoriete_authors_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `authors` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `favoriete_authors`
--

-- Retrieving rows...
LOCK TABLES `favoriete_authors` WRITE;
/*!40000 ALTER TABLE `favoriete_authors` DISABLE KEYS */;
/*!40000 ALTER TABLE `favoriete_authors` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table favoriete_books...

--
-- Table structure for table `favoriete_books`
--

DROP TABLE IF EXISTS `favoriete_books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `favoriete_books` (
  `userid` int NOT NULL,
  `book_id` int NOT NULL,
  KEY `book_id` (`book_id`),
  KEY `userid` (`userid`),
  CONSTRAINT `favoriete_books_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`),
  CONSTRAINT `favoriete_books_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `favoriete_books`
--

-- Retrieving rows...
LOCK TABLES `favoriete_books` WRITE;
/*!40000 ALTER TABLE `favoriete_books` DISABLE KEYS */;
/*!40000 ALTER TABLE `favoriete_books` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table roles...

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `role_id` int NOT NULL,
  `user_role` varchar(100) NOT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `roles`
--

-- Retrieving rows...
LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin'),(2,'user');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table schema_migrations...

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `schema_migrations`
--

-- Retrieving rows...
LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (1,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;
-- Retrieving table structure for table users...

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `userid` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `usersurname` varchar(100) NOT NULL,
  `userpatrynomic` varchar(100) DEFAULT NULL,
  `userphone` varchar(100) NOT NULL,
  `useremail` varchar(100) NOT NULL,
  `userhash` varchar(100) NOT NULL,
  `userrole` int NOT NULL,
  PRIMARY KEY (`userid`),
  KEY `userrole` (`userrole`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`userrole`) REFERENCES `roles` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
-- Sending SELECT query...

--
-- Dumping data for table `users`
--

-- Retrieving rows...
LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
-- Disconnecting from db...
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-06 17:20:43
