install go from the link below:
https://go.dev/doc/install

install MYSQL workbench:
https://dev.mysql.com/downloads/workbench/

run the below queries using the database name as epharmacy:

use epharmacy;
DROP TABLE IF EXISTS `plist`;
CREATE TABLE `plist` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(30) NOT NULL,
    `company` varchar(30) NOT NULL,
	`category` varchar(30) NOT NULL,
    `description` varchar(30) NOT NULL,
    `url` varchar(300) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

use epharmacy;
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
    `id` int(6) unsigned NOT NULL,
    `name` varchar(30) NOT NULL,
    `company` varchar(30) NOT NULL,
	`category` varchar(30) NOT NULL,
    `description` varchar(30) NOT NULL,
    `url` varchar(300) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=latin1;

use epharmacy;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(30) NOT NULL,
    `password` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

use epharmacy;
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(30) NOT NULL,
    `password` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

use epharmacy;
INSERT INTO `user` (`id`, `username`, `password`) VALUES
(1, 'user', 'user');
INSERT INTO `admin` (`id`, `username`, `password`) VALUES
(1, 'admin', 'admin');

use epharmacy;
DROP TABLE IF EXISTS `upload`;
CREATE TABLE `upload` (
    `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
    `prescription` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

open the terminal at the location of your application (you downloaded from git) and run the command:
go run main.go

Tadaaa! your epharmacy application is running on the specified port.