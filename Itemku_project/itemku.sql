/*
SQLyog Community v13.1.8 (64 bit)
MySQL - 10.4.22-MariaDB : Database - itemku
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`itemku` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `itemku`;

/*Table structure for table `detail_transaksis` */

DROP TABLE IF EXISTS `detail_transaksis`;

CREATE TABLE `detail_transaksis` (
  `transaksi_id` bigint(20) DEFAULT NULL,
  `jumlah_item` bigint(20) DEFAULT NULL,
  `harga_transaksi` bigint(20) DEFAULT NULL,
  `keterangan` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `detail_transaksis` */

insert  into `detail_transaksis`(`transaksi_id`,`jumlah_item`,`harga_transaksi`,`keterangan`) values 
(1,100,75000,'selesai');

/*Table structure for table `games` */

DROP TABLE IF EXISTS `games`;

CREATE TABLE `games` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nama` longtext DEFAULT NULL,
  `jenis_item` longtext DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `games` */

insert  into `games`(`id`,`nama`,`jenis_item`) values 
(1,'Mobile Legends','Diamond'),
(2,'Free Fire','Diamond'),
(3,'CODM','CP'),
(4,'PUBG Mobile','UC');

/*Table structure for table `toko_details` */

DROP TABLE IF EXISTS `toko_details`;

CREATE TABLE `toko_details` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `toko_id` bigint(20) DEFAULT NULL,
  `game_id` bigint(20) DEFAULT NULL,
  `jumlah_item` bigint(20) DEFAULT NULL,
  `harga_jual` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

/*Data for the table `toko_details` */

insert  into `toko_details`(`id`,`toko_id`,`game_id`,`jumlah_item`,`harga_jual`) values 
(1,1,1,100,75000),
(2,1,1,75,60000),
(3,1,1,50,40000),
(4,1,1,30,20000),
(5,2,1,90,70000),
(6,2,1,150,150000);

/*Table structure for table `tokos` */

DROP TABLE IF EXISTS `tokos`;

CREATE TABLE `tokos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `users_id` bigint(20) DEFAULT NULL,
  `nama_toko` longtext DEFAULT NULL,
  `create_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `tokos` */

insert  into `tokos`(`id`,`users_id`,`nama_toko`,`create_at`,`update_at`) values 
(1,1,'Toko Zyoler','2022-06-27 11:15:03.000','2022-06-27 11:15:05.000'),
(2,2,'Toko Kusuma Bakti','2022-06-27 11:15:29.000','2022-06-27 11:15:31.000');

/*Table structure for table `transaksis` */

DROP TABLE IF EXISTS `transaksis`;

CREATE TABLE `transaksis` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `game_id` bigint(20) DEFAULT NULL,
  `users_id` bigint(20) DEFAULT NULL,
  `toko_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

/*Data for the table `transaksis` */

insert  into `transaksis`(`id`,`game_id`,`users_id`,`toko_id`) values 
(1,1,3,1);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nama` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `create_at` datetime(3) DEFAULT NULL,
  `update_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `users` */

insert  into `users`(`id`,`nama`,`email`,`password`,`create_at`,`update_at`) values 
(1,'Dani Hidayat','dani@gmail.com','12345','2022-06-27 11:11:40.000','2022-06-27 11:11:45.000'),
(2,'Romi Kusuma Bakti','romi@gmail.com','12345','2022-06-27 11:12:19.000','2022-06-27 11:12:22.000'),
(3,'Jamil Hamdi Harahap','jamil@gmail.com','12345','2022-06-27 11:12:55.000','2022-06-27 11:12:57.000');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
