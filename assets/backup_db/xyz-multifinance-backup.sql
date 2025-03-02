-- MySQL dump 10.13  Distrib 9.2.0, for macos15.2 (arm64)
--
-- Host: localhost    Database: xyz_multifinance
-- ------------------------------------------------------
-- Server version	9.2.0

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
-- Table structure for table `consumers`
--

DROP TABLE IF EXISTS `consumers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consumers` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nik` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `full_name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `legal_name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `place_of_birth` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `date_of_birth` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `salary` double NOT NULL,
  `ktp_image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `selfie_image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_verified` tinyint(1) DEFAULT '0',
  `rejection_reason` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nix_nik` (`nik`),
  KEY `idx_consumers_deleted_at` (`deleted_at`),
  KEY `fk_users_consumer` (`user_id`),
  CONSTRAINT `fk_users_consumer` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consumers`
--

LOCK TABLES `consumers` WRITE;
/*!40000 ALTER TABLE `consumers` DISABLE KEYS */;
INSERT INTO `consumers` VALUES ('00d0c9e2-023b-433d-8bde-38031442123c',NULL,'','','','','',0,'','',0,'','2025-03-02 15:41:23','2025-03-02 15:41:23',NULL,'5c24352a-1b25-4002-a160-41788a00a783'),('01a7da8e-f22a-41f1-b996-4231938ad7c0',NULL,'','','','','',0,'','',0,'','2025-03-02 15:44:55','2025-03-02 15:44:55',NULL,'43ed5293-9371-4bac-9e5a-c04411b2a79a'),('8f62cb1a-613b-4a68-8fee-058c544843fb',NULL,'','','','','',0,'','',0,'','2025-03-02 15:41:12','2025-03-02 15:41:12',NULL,'ae578a67-1063-4f11-88eb-ae3d7186e7b1'),('8ffb36d1-234b-4adf-a213-fb52699fb581','1234567890123459','Adelia Risky Santoso','Adelia','087890987890','Sumenep','01-10-2001',10000000,'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTpR24fRQKQUeb_vZ5mhvekQZvj5iSfiFhKNw&s','https://i.pinimg.com/736x/50/08/ef/5008efb9df96969624d2674645027a3a.jpg',1,'Data diri anda belum lengkap','2025-03-02 15:39:48','2025-03-02 16:02:24',NULL,'30029a52-d79d-41e3-8279-97243f30863a');
/*!40000 ALTER TABLE `consumers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `limits`
--

DROP TABLE IF EXISTS `limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `limits` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `limit_available` double NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `consumer_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_limits_deleted_at` (`deleted_at`),
  KEY `fk_consumers_limits` (`consumer_id`),
  CONSTRAINT `fk_consumers_limits` FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `limits`
--

LOCK TABLES `limits` WRITE;
/*!40000 ALTER TABLE `limits` DISABLE KEYS */;
INSERT INTO `limits` VALUES ('f2b33782-8b42-40ab-920f-abfaa96d21db',41775041.56,'2025-03-02 16:02:23','fa1343ac-d7aa-4d99-8dd3-3603a40de5a2','2025-03-02 16:53:54',NULL,NULL,NULL,'8ffb36d1-234b-4adf-a213-fb52699fb581');
/*!40000 ALTER TABLE `limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` datetime(3) NOT NULL,
  `amount_paid` double NOT NULL,
  `status` enum('PENDING','SUCCESS','FAILED') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `transaction_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_payments_deleted_at` (`deleted_at`),
  KEY `fk_transactions_payments` (`transaction_id`),
  CONSTRAINT `fk_transactions_payments` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES ('020cadcd-d663-461b-a767-f6ff20c5eadf','2025-03-03 02:19:45.220',351663.37,'SUCCESS','2025-03-02 19:19:45','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:45',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('0cc683d7-bfb9-4e2b-a407-130bcaa7662a','2025-03-03 02:21:44.001',351663.37,'SUCCESS','2025-03-02 19:21:44','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:21:44',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('2cf5ce84-f9aa-402a-b29d-a4182fac4dab','2025-03-03 02:20:44.466',351663.37,'SUCCESS','2025-03-02 19:20:44','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:20:44',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('2d47a196-1430-47ce-aca0-77554015d09a','2025-03-03 02:19:50.084',351663.37,'SUCCESS','2025-03-02 19:19:50','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:50',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('609f8ab9-52ab-4ca3-b7e2-4cc1f5bc21fb','2025-03-03 02:19:46.627',351663.37,'SUCCESS','2025-03-02 19:19:46','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:46',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('71d15a7e-f796-480e-8db4-53ad45e75da3','2025-03-03 02:19:48.573',351663.37,'SUCCESS','2025-03-02 19:19:48','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:48',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('7560435e-a3d9-419d-8199-6ce2e881d112','2025-03-03 02:19:44.117',351663.37,'SUCCESS','2025-03-02 19:19:44','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:44',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('79879450-c72d-4556-9cf1-740be9e41cb1','2025-03-03 02:19:51.863',351663.37,'SUCCESS','2025-03-02 19:19:51','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:19:51',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('7ba16066-4981-4dd1-a197-569a955ca57f','2025-03-03 02:16:49.196',351663.37,'SUCCESS','2025-03-02 19:16:49','5c24352a-1b25-4002-a160-41788a00a783','2025-03-02 19:16:49',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('bba48e6f-c1b9-45b4-a67f-cbf0eb6d0813','2025-03-03 02:21:43.299',351663.37,'SUCCESS','2025-03-02 19:21:43','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:21:43',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('c028cd5f-6480-42bd-81d6-b92c4ed891b9','2025-03-03 02:21:42.419',351663.37,'SUCCESS','2025-03-02 19:21:42','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:21:42',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15'),('ca4cebc7-57d4-4141-a545-d3bb57adab8a','2025-03-03 02:21:44.744',351663.37,'SUCCESS','2025-03-02 19:21:44','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 19:21:44',NULL,NULL,NULL,'0812692a-e129-4d00-be67-8b42fa4d0b15');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES ('87ee65c2-e978-45dd-a071-14bfc6a9669c','Admin','2025-03-02 15:39:48','2025-03-02 15:39:48',NULL),('ac704435-f8aa-4d45-94ca-2d2e7133fe5f','User','2025-03-02 15:39:48','2025-03-02 15:39:48',NULL);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `contract_number` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `otr` double NOT NULL,
  `tenor` bigint NOT NULL,
  `admin_fee` double NOT NULL,
  `installment_amt` double NOT NULL,
  `amount_interest` double NOT NULL,
  `asset_name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `consumer_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_transactions_contract_number` (`contract_number`),
  KEY `idx_transactions_deleted_at` (`deleted_at`),
  KEY `fk_consumers_transaction` (`consumer_id`),
  CONSTRAINT `fk_consumers_transaction` FOREIGN KEY (`consumer_id`) REFERENCES `consumers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('0812692a-e129-4d00-be67-8b42fa4d0b15','CN-20250302-L3TD81',3999998,12,5000,351663.37,219962.44,'Airpods 2','2025-03-02 16:53:54','30029a52-d79d-41e3-8279-97243f30863a','2025-03-02 16:53:54',NULL,NULL,NULL,'8ffb36d1-234b-4adf-a213-fb52699fb581');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_id` char(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `fk_roles_users` (`role_id`),
  CONSTRAINT `fk_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('30029a52-d79d-41e3-8279-97243f30863a','user@example.com','$2a$10$znBHpAJz1uVX18Io1fd4yOhVIYfKSjtvG9xb.o3ps9WjmLUu8.Ec2','2025-03-02 15:39:48','2025-03-02 15:39:48',NULL,'ac704435-f8aa-4d45-94ca-2d2e7133fe5f'),('43ed5293-9371-4bac-9e5a-c04411b2a79a','threeuser@example.com','$2a$10$tXlf1OvNmvL9BgZ7jIkLherSN5mVi2pCTAMRTyC4bQUqyv5JuRAqy','2025-03-02 15:44:55','2025-03-02 15:44:55',NULL,'ac704435-f8aa-4d45-94ca-2d2e7133fe5f'),('5c24352a-1b25-4002-a160-41788a00a783','twouser@example.com','$2a$10$r9qEtQ6dbz/wf3362tuh0.zf7Zx2mkNlCNx8to9ChxtQI448oMYc6','2025-03-02 15:41:23','2025-03-02 15:41:23',NULL,'ac704435-f8aa-4d45-94ca-2d2e7133fe5f'),('ae578a67-1063-4f11-88eb-ae3d7186e7b1','oneuser@example.com','$2a$10$gyicXmqjwtDb.Wkao2GqVuo68Go73nARkabZBId/Cy9HY2zNCBnk.','2025-03-02 15:41:12','2025-03-02 15:41:12',NULL,'ac704435-f8aa-4d45-94ca-2d2e7133fe5f'),('fa1343ac-d7aa-4d99-8dd3-3603a40de5a2','admin@example.com','$2a$10$U9B3pieL09DNegqdmfsdP.MvVYVOYobcOv7HFeUbi90kfNiqjfR/u','2025-03-02 15:39:48','2025-03-02 15:39:48',NULL,'87ee65c2-e978-45dd-a071-14bfc6a9669c');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-03  6:28:22
