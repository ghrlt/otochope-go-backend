-- phpMyAdmin SQL Dump
-- version 5.2.3
-- https://www.phpmyadmin.net/
--
-- Host: database
-- Generation Time: Nov 05, 2025 at 03:13 PM
-- Server version: 11.8.3-MariaDB-ubu2404
-- PHP Version: 8.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `otochope`
--

-- --------------------------------------------------------

--
-- Table structure for table `coupons`
--

CREATE TABLE `coupons` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `code` text NOT NULL,
  `value` int(11) NOT NULL,
  `is_value_a_percentage` tinyint(1) NOT NULL DEFAULT 1,
  `total_usage_limit` int(11) DEFAULT NULL,
  `per_user_usage_limit` int(11) DEFAULT NULL,
  `min_cart_total` int(11) DEFAULT NULL,
  `only_valid_after` timestamp NULL DEFAULT NULL,
  `only_valid_before` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `coupons`
--

INSERT INTO `coupons` (`uid`, `code`, `value`, `is_value_a_percentage`, `total_usage_limit`, `per_user_usage_limit`, `min_cart_total`, `only_valid_after`, `only_valid_before`, `created_at`, `active`) VALUES
('857e1b3d-b9cb-11f0-99c4-1ebc133d497c', 'TEST20', 20, 1, 10, 1, 0, '2025-11-01 15:54:21', '2025-11-30 15:54:21', '2025-11-05 14:54:59', 1),
('8c089fd8-b9cb-11f0-99c4-1ebc133d497c', 'TEST30', 30, 0, NULL, NULL, NULL, NULL, NULL, '2025-11-05 14:55:10', 1);

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `category_uid` uuid NOT NULL,
  `label` text NOT NULL,
  `description` text NOT NULL,
  `price_in_eur_cents` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`uid`, `category_uid`, `label`, `description`, `price_in_eur_cents`, `created_at`, `active`) VALUES
('56ca7b59-b9cb-11f0-99c4-1ebc133d497c', '29a572f2-b9cb-11f0-99c4-1ebc133d497c', 'TEST ITEM', 'item is a test item, member of TEST CATEGORY', 500, '2025-11-05 14:53:40', 1),
('65d4a360-b9cb-11f0-99c4-1ebc133d497c', '3d298d65-b9cb-11f0-99c4-1ebc133d497c', 'TEST ITEM #2', 'item is member of a child category', 100, '2025-11-05 14:54:05', 1);

-- --------------------------------------------------------

--
-- Table structure for table `item_categories`
--

CREATE TABLE `item_categories` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `parent_category_uid` uuid DEFAULT NULL,
  `label` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `item_categories`
--

INSERT INTO `item_categories` (`uid`, `parent_category_uid`, `label`, `created_at`, `active`) VALUES
('29a572f2-b9cb-11f0-99c4-1ebc133d497c', NULL, 'TEST CATEGORY', '2025-11-05 14:52:24', 1),
('3286a601-b9cb-11f0-99c4-1ebc133d497c', NULL, 'TEST PARENT CATEGORY', '2025-11-05 14:52:39', 1),
('3d298d65-b9cb-11f0-99c4-1ebc133d497c', '3286a601-b9cb-11f0-99c4-1ebc133d497c', 'TEST CHILD CATEGORY', '2025-11-05 14:52:57', 1);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `user_uid` uuid NOT NULL,
  `cart_uid` uuid NOT NULL,
  `amount_in_eur_cents` int(11) NOT NULL,
  `amount` text NOT NULL,
  `currency` smallint(6) NOT NULL,
  `status` varchar(50) NOT NULL,
  `reference` text NOT NULL,
  `charge_reference` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `telegram_identifier` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`uid`, `telegram_identifier`, `created_at`, `active`) VALUES
('b2ad29ca-b912-11f0-99c4-1ebc133d497c', 0, '2025-11-04 00:10:40', 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_carts`
--

CREATE TABLE `user_carts` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `user_uid` uuid NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `user_carts`
--

INSERT INTO `user_carts` (`uid`, `user_uid`, `created_at`, `updated_at`, `active`) VALUES
('98652b2d-b9cb-11f0-99c4-1ebc133d497c', 'b2ad29ca-b912-11f0-99c4-1ebc133d497c', '2025-11-05 14:55:30', '2025-11-05 14:55:30', 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_cart_coupons`
--

CREATE TABLE `user_cart_coupons` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `cart_uid` uuid NOT NULL,
  `coupon_uid` uuid NOT NULL,
  `added_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `user_cart_coupons`
--

INSERT INTO `user_cart_coupons` (`uid`, `cart_uid`, `coupon_uid`, `added_at`, `active`) VALUES
('677f1bfb-b9cc-11f0-99c4-1ebc133d497c', '98652b2d-b9cb-11f0-99c4-1ebc133d497c', '8c089fd8-b9cb-11f0-99c4-1ebc133d497c', '2025-11-05 15:01:18', 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_cart_items`
--

CREATE TABLE `user_cart_items` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `cart_uid` uuid NOT NULL,
  `item_uid` uuid NOT NULL,
  `quantity` smallint(6) NOT NULL,
  `added_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `user_cart_items`
--

INSERT INTO `user_cart_items` (`uid`, `cart_uid`, `item_uid`, `quantity`, `added_at`, `updated_at`, `active`) VALUES
('9c62eba1-b9cc-11f0-99c4-1ebc133d497c', '98652b2d-b9cb-11f0-99c4-1ebc133d497c', '56ca7b59-b9cb-11f0-99c4-1ebc133d497c', 10, '2025-11-05 15:02:46', '2025-11-05 15:02:46', 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_inventory`
--

CREATE TABLE `user_inventory` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `user_uid` uuid NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `user_inventory`
--

INSERT INTO `user_inventory` (`uid`, `user_uid`, `active`) VALUES
('16eb4011-b9cb-11f0-99c4-1ebc133d497c', 'b2ad29ca-b912-11f0-99c4-1ebc133d497c', 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_inventory_items`
--

CREATE TABLE `user_inventory_items` (
  `uid` uuid NOT NULL DEFAULT uuid(),
  `inventory_uid` uuid NOT NULL,
  `item_uid` uuid NOT NULL,
  `quantity` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `active` tinyint(1) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dumping data for table `user_inventory_items`
--

INSERT INTO `user_inventory_items` (`uid`, `inventory_uid`, `item_uid`, `quantity`, `created_at`, `updated_at`, `active`) VALUES
('af30448b-b9cb-11f0-99c4-1ebc133d497c', '16eb4011-b9cb-11f0-99c4-1ebc133d497c', '56ca7b59-b9cb-11f0-99c4-1ebc133d497c', 2, '2025-11-05 14:56:09', NULL, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `coupons`
--
ALTER TABLE `coupons`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `item_categories`
--
ALTER TABLE `item_categories`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`uid`),
  ADD UNIQUE KEY `unique_tlg_id` (`telegram_identifier`);

--
-- Indexes for table `user_carts`
--
ALTER TABLE `user_carts`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `user_cart_items`
--
ALTER TABLE `user_cart_items`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `user_inventory`
--
ALTER TABLE `user_inventory`
  ADD PRIMARY KEY (`uid`);

--
-- Indexes for table `user_inventory_items`
--
ALTER TABLE `user_inventory_items`
  ADD PRIMARY KEY (`uid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
