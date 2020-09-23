SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `treemap`;
CREATE DATABASE `treemap` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `treemap`;

DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL,
                        `node_type` tinyint unsigned NOT NULL,
                        `description` text COLLATE utf8mb4_general_ci NOT NULL,
                        `children` json NOT NULL,
                        `leaf_type` tinyint unsigned NOT NULL,
                        `content` text COLLATE utf8mb4_general_ci NOT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `node` (`id`, `name`, `node_type`, `description`, `children`, `leaf_type`, `content`) VALUES
(1,	'WEB安全',	0,	'',	'[]',	0,	''),
(2,	'test2',	0,	'',	'[]',	0,	'');
