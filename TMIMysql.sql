CREATE TABLE `user` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `password` varchar(255),
  `nomorHP` integer,
  `email` varchar(255)
);

CREATE TABLE `mahasiswa` (
  `id` integer PRIMARY KEY,
  `url_foto` varchar(255),
  `nama_lengkap` varchar(255),
  `tanggal_lahir` varchar(255),
  `jenis_kelamin` varchar(255),
  `asal_kampus` varchar(255),
  `nim` varchar(255),
  `jurusan` varchar(255),
  `tahun_masuk` varchar(255),
  `kota_kabupaten` varchar(255),
  `id_tantangan` integer,
  `id_silabus` integer,
  `id_event` integer,
  `id_user` integer
);

CREATE TABLE `tantangan` (
  `id` integer PRIMARY KEY,
  `judul_tantangan` varchar(255),
  `tema` varchar(255),
  `latar` varchar(255),
  `url_video_tantangan` varchar(255),
  `task_tantangan` varchar(255),
  `ujian_tantangan` varchar(255),
  `skor_tantangan` varchar(255),
  `id_admin` integer
);

CREATE TABLE `silabus` (
  `id` integer PRIMARY KEY,
  `judul_silabus` varchar(255),
  `definisi` varchar(255),
  `fungsi_silabus` varchar(255),
  `deskripsi` varchar(255),
  `url_video_silabus` varchar(255),
  `task_silabus` varchar(255),
  `ujian_silabus` varchar(255),
  `id_admin` integer
);

CREATE TABLE `admin` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `password` varchar(255),
  `role` varchar(255)
);

CREATE TABLE `event` (
  `id` integer PRIMARY KEY,
  `judul_event` varchar(255),
  `deskripsi_event` varchar(255),
  `kriteria_event` varchar(255),
  `tanggal_event` varchar(255),
  `id_admin` integer
);

CREATE TABLE `leaderboard` (
  `ranking` integer,
  `id_mahasiswa` integer,
  `id_tantangan` integer,
  `id_silabus` integer
);

CREATE TABLE `statusSilabus` (
  `skor_silabus` integer,
  `id_mahasiswa` integer,
  `id_silabus` integer,
  `checkSilabus` integer
);

CREATE TABLE `statusTantangan` (
  `skor_tantangan` integer,
  `id_mahasiswa` integer,
  `id_tantangan` integer,
  `checktantangan` integer
);

CREATE TABLE `statusEvent` (
  `checktantangan` integer,
  `id_mahasiswa` integer,
  `id_event` integer
);

ALTER TABLE `tantangan` ADD FOREIGN KEY (`id`) REFERENCES `mahasiswa` (`id_tantangan`);

ALTER TABLE `silabus` ADD FOREIGN KEY (`id`) REFERENCES `mahasiswa` (`id_silabus`);

ALTER TABLE `event` ADD FOREIGN KEY (`id`) REFERENCES `mahasiswa` (`id_event`);

ALTER TABLE `mahasiswa` ADD FOREIGN KEY (`id_user`) REFERENCES `user` (`id`);

ALTER TABLE `tantangan` ADD FOREIGN KEY (`id_admin`) REFERENCES `admin` (`id`);

ALTER TABLE `silabus` ADD FOREIGN KEY (`id_admin`) REFERENCES `admin` (`id`);

ALTER TABLE `event` ADD FOREIGN KEY (`id_admin`) REFERENCES `admin` (`id`);

ALTER TABLE `leaderboard` ADD FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id`);

ALTER TABLE `leaderboard` ADD FOREIGN KEY (`id_tantangan`) REFERENCES `tantangan` (`id`);

ALTER TABLE `leaderboard` ADD FOREIGN KEY (`id_silabus`) REFERENCES `silabus` (`id`);

ALTER TABLE `statusSilabus` ADD FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id`);

ALTER TABLE `statusSilabus` ADD FOREIGN KEY (`id_silabus`) REFERENCES `silabus` (`id`);

ALTER TABLE `statusTantangan` ADD FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id`);

ALTER TABLE `statusTantangan` ADD FOREIGN KEY (`id_tantangan`) REFERENCES `tantangan` (`id`);

ALTER TABLE `statusEvent` ADD FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id`);

ALTER TABLE `statusEvent` ADD FOREIGN KEY (`id_event`) REFERENCES `event` (`id`);
