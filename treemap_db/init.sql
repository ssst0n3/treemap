-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `index` int NOT NULL,
  `name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `node_type` tinyint unsigned NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `leaf_type` tinyint unsigned NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `node` (`id`, `index`, `name`, `node_type`, `description`, `leaf_type`, `content`) VALUES
(1,	0,	'WEB安全',	0,	'',	0,	''),
(2,	0,	'test2',	0,	'',	0,	''),
(6,	0,	'xxxxxxx',	0,	'',	0,	''),
(7,	0,	'',	0,	'',	0,	''),
(8,	0,	'',	0,	'',	0,	''),
(9,	0,	'test1',	1,	'',	0,	'');

DROP TABLE IF EXISTS `node_relation`;
CREATE TABLE `node_relation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent` int unsigned NOT NULL,
  `child` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `node_relation` (`id`, `parent`, `child`) VALUES
(1,	1,	3),
(2,	1,	4),
(3,	1,	5),
(4,	3,	6),
(5,	4,	7),
(6,	4,	8),
(7,	1,	9);

-- 2020-10-06 15:34:16