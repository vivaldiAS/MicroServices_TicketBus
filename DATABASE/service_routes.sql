-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 15 Bulan Mei 2024 pada 18.12
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `service_routes`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `routes`
--

CREATE TABLE `routes` (
  `id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `derpature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `arrival` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `harga` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `brand_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `routes`
--

INSERT INTO `routes` (`id`, `derpature`, `arrival`, `harga`, `type`, `status`, `brand_id`, `created_at`, `updated_at`) VALUES
(1, 'Medan', 'Penyabungan', '1', 'Ekonomi', 1, 1, '2023-05-12 20:42:58', '2024-01-15 02:25:02'),
(2, 'Medan', 'Tantom', '170000', 'Ekonomi', 0, 1, '2023-05-12 20:44:33', '2023-05-23 22:04:21'),
(3, 'Medan', 'P.Sidimpuan', '160000', 'Ekonomi', 1, 1, '2023-05-12 20:45:11', '2023-05-12 20:45:11'),
(4, 'Medan', 'Sipirok', '150000', 'Ekonomi', 1, 1, '2023-05-12 20:45:48', '2023-05-12 20:45:48'),
(5, 'Medan', 'Tarutung', '140000', 'Eksekutif', 1, 1, '2023-05-12 20:46:19', '2023-05-12 20:46:19'),
(6, 'Medan', 'Sarulla', '120000', 'Ekonomi', 1, 1, '2023-05-12 20:46:55', '2023-05-12 20:46:55'),
(7, 'Medan', 'Aek raja', '115000', 'Ekonomi', 1, 1, '2023-05-12 20:47:34', '2023-05-12 20:47:34'),
(8, 'Medan', 'O.Hasang', '155000', 'Ekonomi', 1, 1, '2023-05-12 20:48:49', '2023-05-12 20:48:49'),
(9, 'Medan', 'O.Tukka', '115000', 'Ekonomi', 1, 1, '2023-05-12 20:49:43', '2023-05-12 20:49:43'),
(10, 'Medan', 'Bakkara', '110000', 'Ekonomi', 1, 1, '2023-05-12 20:52:42', '2023-05-12 20:52:42'),
(11, 'Medan', 'Lobuksikkam', '110000', 'Ekonomi', 1, 1, '2023-05-12 20:53:13', '2023-05-12 20:53:13'),
(12, 'Medan', 'Simamora Nabolak', '110000', 'Ekonomi', 1, 1, '2023-05-12 20:53:55', '2023-05-12 20:53:55'),
(13, 'Medan', 'Sipahutar', '110000', 'Ekonomi', 1, 1, '2023-05-12 20:54:45', '2023-05-12 20:54:45'),
(14, 'Medan', 'Tarutung', '110000', 'Ekonomi', 1, 1, '2023-05-12 20:55:12', '2023-05-12 20:55:12'),
(15, 'Medan', 'Muara', '105000', 'Ekonomi', 1, 1, '2023-05-12 20:55:53', '2023-05-12 20:55:53'),
(16, 'Medan', 'Siborong-borong', '95000', 'Ekonomi', 1, 1, '2023-05-12 20:56:47', '2023-05-12 20:56:47'),
(17, 'Medan', 'Toba', '85000', 'Ekonomi', 1, 1, '2023-05-12 20:57:26', '2023-05-12 20:57:26'),
(18, 'Medan', 'Parapat', '80000', 'Ekonomi', 1, 1, '2023-05-12 20:58:01', '2023-05-12 20:58:01'),
(19, 'Medan', 'P.Siantar', '65000', 'Ekonomi', 1, 1, '2023-05-12 20:58:27', '2023-05-12 20:58:27'),
(20, 'Bandara Kualanamu', 'Tarutung', '180000', 'Eksekutif', 1, 1, '2023-05-12 20:59:13', '2023-05-12 20:59:13'),
(21, 'P.Siantar', 'Tarutung', '85000', 'Ekonomi', 1, 1, '2023-05-12 20:59:59', '2023-05-12 20:59:59'),
(22, 'P.Siantar', 'Bakkara', '85000', 'Ekonomi', 1, 1, '2023-05-12 21:00:31', '2023-05-12 21:00:31'),
(23, 'P.Siantar', 'Toba', '65000', 'Ekonomi', 1, 1, '2023-05-12 21:00:56', '2023-05-12 21:00:56'),
(24, 'P.Siantar', 'Siborong-borong', '75000', 'Ekonomi', 1, 1, '2023-05-12 21:01:26', '2023-05-12 21:01:26'),
(25, 'Parapat', 'Tarutung', '75000', 'Ekonomi', 1, 1, '2023-05-12 21:01:51', '2023-05-12 21:01:51'),
(26, 'Parapat', 'Bakkara', '75000', 'Ekonomi', 1, 1, '2023-05-12 21:02:10', '2023-05-12 21:02:10'),
(27, 'P.Sidimpuan', 'Tarutung', '80000', 'Ekonomi', 1, 1, '2023-05-12 21:02:46', '2023-05-12 21:02:46'),
(28, 'Tebing Tinggi', 'Tarutung', '95000', 'Ekonomi', 1, 1, '2023-05-12 21:03:13', '2023-05-12 21:03:13'),
(29, 'Tebing Tinggi', 'Bakkara', '95000', 'Ekonomi', 1, 1, '2023-05-12 21:03:35', '2023-05-12 21:03:35'),
(30, 'Tarutung', 'Medan', '140000', 'Eksekutif', 1, 1, '2023-07-16 18:58:01', '2023-07-16 18:58:01'),
(31, 'Medan', 'Balige', '80000', 'Ekonomi', 1, 1, '2023-07-18 00:48:12', '2023-07-18 00:48:12'),
(32, 'Medan', 'Jakarta', '10000', 'Ekonomi', 1, 2, '2024-05-10 03:01:37', '2024-05-12 03:20:11'),
(33, 'Sidikalang', 'Kabanjahe', '25000', 'Ekonomi', 1, 2, '2024-05-10 04:50:46', '2024-05-11 21:15:29');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
