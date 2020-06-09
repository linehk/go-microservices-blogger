CREATE DATABASE IF NOT EXISTS `blog` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `blog`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0',
  `title` varchar(100) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  `content` text,
  `created_on` int(10) unsigned DEFAULT '0',
  `created_by` varchar(100) DEFAULT '',
  `modified_on` int(10) unsigned DEFAULT '0',
  `modified_by` varchar(255) DEFAULT '',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `created_on` int(10) unsigned DEFAULT '0',
  `created_by` varchar(100) DEFAULT '',
  `modified_on` int(10) unsigned DEFAULT '0',
  `modified_by` varchar(100) DEFAULT '',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;