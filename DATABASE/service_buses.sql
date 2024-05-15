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
-- Database: `service_buses`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `buses`
--

CREATE TABLE `buses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `supir_id` bigint(20) UNSIGNED NOT NULL,
  `loket_id` bigint(20) UNSIGNED NOT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `police_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `number_of_seats` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `merk_id` bigint(20) UNSIGNED NOT NULL,
  `nomor_pintu` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `buses`
--

INSERT INTO `buses` (`id`, `supir_id`, `loket_id`, `type`, `police_number`, `number_of_seats`, `merk_id`, `nomor_pintu`, `status`, `created_at`, `updated_at`) VALUES
(1, 6, 1, 'Ekonomi', 'BB 0938 KB', '12', 1, 'KBT 001', 1, '2023-05-12 21:20:00', '2023-05-16 20:26:22'),
(2, 5, 1, 'Eksekutif', 'BB 1234 KM', '13', 1, 'KBT 002', 1, '2023-05-16 01:35:07', '2023-05-16 02:06:31'),
(3, 7, 1, 'Ekonomi', 'BK 0978 MK', '12', 1, 'KBT 003', 1, '2023-05-16 02:02:05', '2023-05-22 00:27:48'),
(4, 15, 2, 'Ekonomi', 'BK 1232 KM', '12', 1, 'KBT 004', 0, '2023-05-23 20:09:49', '2023-08-13 20:55:46'),
(5, 16, 2, 'Eksekutif', 'BB 1234 KF', '13', 1, 'KBT EKS 001', 1, '2023-05-23 22:05:48', '2023-05-23 22:05:48'),
(6, 18, 2, 'Ekonomi', 'BB 5555 FH', '12', 1, 'KBT 005', 1, '2023-07-12 19:40:43', '2023-07-12 19:40:43'),
(7, 22, 2, 'Ekonomi', 'BK 2340 KM', '12', 1, 'KBT 010', 1, '2023-07-18 00:38:51', '2023-07-18 00:38:51'),
(8, 28, 2, 'Ekonomi', 'BK 4321 KL', '12', 1, 'KBT 999', 1, '2023-08-07 19:14:16', '2023-08-07 19:14:16'),
(9, 14, 5, 'Ekonomi', 'BB 2351 LD', '10', 2, '69', 1, '2024-05-09 09:29:30', '2024-05-12 03:18:45'),
(10, 26, 5, 'Ekonomi', 'ABC123', '12', 2, '12', 1, '2024-05-09 11:30:11', '2024-05-12 03:18:56'),
(18, 35, 4, 'Ekonomi', 'bb 999cc', '12', 2, 'TTI 00', 1, '2024-05-12 03:12:15', '2024-05-12 03:12:15'),
(19, 5, 4, 'Ekonomi', 'bb 000 cc', '12', 2, 'TTI 55', 1, '2024-05-12 03:13:36', '2024-05-12 03:13:36'),
(20, 5, 1, 'AC', 'B 1234 ABC', '30', 0, '2', 1, NULL, NULL),
(22, 5, 1, 'AC', 'B 1234 ABC', '30', 0, '2', 1, NULL, NULL),
(23, 8, 2, 'Ekonomi', 'bk 999 cc', '12', 0, 'KBT 77', 1, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `buses`
--
ALTER TABLE `buses`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `buses`
--
ALTER TABLE `buses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
