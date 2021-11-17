-- name: ViewMahasiswa :many
SELECT * FROM mahasiswa;

-- name: CreateMahasiswa :exec
INSERT INTO mahasiswa(
    username, password, email
) VALUES (
    ?,?,?
);

-- name: GetMahasiswaByAuth :one
SELECT * FROM mahasiswa WHERE username = ? AND password = ? OR email = ? AND password = ?;

-- name: GetMahasiswa :one
SELECT * FROM mahasiswa WHERE id =?;

-- name: UpdateMahasiswa :exec
UPDATE mahasiswa SET email = ?, url_foto = ?, nama_lengkap = ?, tanggal_lahir = ?, jenis_kelamin = ?, asal_kampus = ? , nim = ?, jurusan = ?, tahun_masuk = ?, kota_kabupaten = ? WHERE id = ?; 