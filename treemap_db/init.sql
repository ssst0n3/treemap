-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `index` int unsigned NOT NULL,
  `name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `node_type` tinyint unsigned NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `content_id` int unsigned NOT NULL,
  `content_type` int unsigned NOT NULL,
  `leaf_type` tinyint unsigned NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `node` (`id`, `index`, `name`, `node_type`, `description`, `content_id`, `content_type`, `leaf_type`, `content`) VALUES
(1,	0,	'WEB安全',	0,	'',	0,	0,	0,	''),
(17,	1,	'a',	1,	'',	0,	0,	0,	''),
(18,	3,	'b',	1,	'',	0,	0,	0,	''),
(19,	7,	'c',	1,	'',	0,	0,	0,	''),
(20,	6,	'd',	1,	'',	0,	0,	0,	'');

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
(7,	1,	9),
(8,	1,	10),
(9,	1,	11),
(10,	10,	12),
(11,	1,	13),
(12,	0,	14),
(13,	1,	15),
(14,	1,	16),
(15,	1,	17),
(16,	1,	18),
(17,	1,	19),
(18,	1,	20);

-- 2020-10-08 09:36:05