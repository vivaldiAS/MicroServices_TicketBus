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
-- Database: `service_lokets`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `lokets`
--

CREATE TABLE `lokets` (
  `id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `nama_loket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `lokasi_loket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `admin_id` bigint(20) UNSIGNED NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `brand_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `lokets`
--

INSERT INTO `lokets` (`id`, `nama_loket`, `lokasi_loket`, `admin_id`, `status`, `brand_id`, `created_at`, `updated_at`) VALUES
(1, 'Loket Medan', 'Medan', 3, 1, 1, '2023-05-11 21:47:06', '2024-05-10 01:01:39'),
(2, 'Loket Tarutung', 'Tarutung Jalan DI.Panjaitan no.27 lt.1', 8, 1, 1, '2023-05-11 23:45:31', '2023-07-16 18:56:25'),
(3, 'Loket Siborong-borong', 'Jln.Gereja Siborong-borong', 9, 1, 1, '2023-07-18 00:18:51', '2023-07-18 00:18:51'),
(4, 'Loket TTI Sidikalang Pedalaman', 'Sidikalang', 34, 1, 2, NULL, '2024-05-12 00:18:11'),
(5, 'Loket TTI Jakarta', 'Jakarta', 37, 1, 2, NULL, NULL);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
