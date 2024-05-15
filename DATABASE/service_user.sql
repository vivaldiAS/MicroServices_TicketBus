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
-- Database: `service_user`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `roles`
--

INSERT INTO `roles` (`id`, `role`, `created_at`, `updated_at`) VALUES
(1, 'admin_kantor', NULL, NULL),
(2, 'passenger', NULL, NULL),
(3, 'driver', NULL, NULL),
(4, 'admin_loket', NULL, NULL),
(5, 'direksi', NULL, NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `photo` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `password` varchar(255) NOT NULL,
  `role_id` varchar(100) NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `phone_number`, `address`, `gender`, `photo`, `status`, `password`, `role_id`, `email_verified_at`, `remember_token`, `created_at`, `updated_at`) VALUES
(1, 'masrin', 'masrin@gmail.com', '081234567890', 'Jl. Bengawan Solo', 'Laki-laki', '', 1, '123456', '2', '2024-05-08 01:55:18', NULL, '2024-05-24 01:57:15', '2024-05-25 01:57:15'),
(2, 'Vivaldi', 'vivaldi@gmail.com', '123456789', 'Jl. Kartini', 'Laki-laki', '', 1, '123456', '1', '2024-05-09 12:23:17', NULL, NULL, NULL),
(4, 'Mark Tuan', 'marktuan@gmail.com', '081234567890', 'Philadelphia', 'Laki-laki', '', 1, '123456', '2', '2024-05-09 12:23:37', NULL, NULL, NULL),
(5, 'Supir1', 'supir1@gmail.com', '88888888', 'Jl. Sibolga', 'Laki-laki', '', 1, '123456', '3', '2024-05-15 00:49:28', NULL, NULL, NULL),
(6, 'Micahel Learns To Rock', 'janedoe@example.com', '081234567891', 'Jl. Contoh No. 124, Jakarta', 'Perempuan', '', 1, '123456', '4', NULL, NULL, NULL, NULL),
(7, 'supir2', 'supir2@gmail.com', '081234567891', 'Jl. Jakarta', 'Perempuan', '', 1, '123456', '3', NULL, NULL, NULL, NULL),
(8, 'Vivaldi', 'vs@gmail.com', '71781', 'Jl. Siantar', 'Laki-laki', '', 1, '123456', '3', NULL, NULL, NULL, NULL),
(9, 'test', 'testi@gmail.com', '73787823', 'test\n', 'Laki-laki', '', 1, '123456', '3', NULL, NULL, NULL, NULL),
(10, 'Jeanet', 'jeaner@gmail.com', '6547382901', 'Jl. Italia', 'Laki-laki', '', 1, '123456', '4', NULL, NULL, NULL, NULL),
(11, 'Backstreet Boys', 'bsboys@gmail.com', '99999999', 'Laguboti', 'Laki-laki', '', 1, '123456', '4', NULL, NULL, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
