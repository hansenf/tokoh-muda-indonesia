-- name: CreateMahasiwa :exec
INSERT INTO mahasiswa (
  id, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, asal_kampus, nim, jurusan, tahun_masuk, kota_kabupaten, id_tantangan, id_silabus, id_event
) VALUES (
  ?,?,?,?,?,?,?,?,?,?,?,?,?
);

