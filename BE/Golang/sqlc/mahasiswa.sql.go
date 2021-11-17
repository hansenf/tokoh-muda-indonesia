// Code generated by sqlc. DO NOT EDIT.
// source: mahasiswa.sql

package db

import (
	"context"
	"database/sql"
)

const createMahasiswa = `-- name: CreateMahasiswa :exec
INSERT INTO mahasiswa(
    username, password, email
) VALUES (
    ?,?,?
)
`

type CreateMahasiswaParams struct {
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
	Email    sql.NullInt32  `json:"email"`
}

func (q *Queries) CreateMahasiswa(ctx context.Context, arg CreateMahasiswaParams) error {
	_, err := q.db.ExecContext(ctx, createMahasiswa, arg.Username, arg.Password, arg.Email)
	return err
}

const getMahasiswa = `-- name: GetMahasiswa :one
SELECT id, username, password, nomorhp, email, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, asal_kampus, nim, jurusan, tahun_masuk, kota_kabupaten, id_tantangan, id_silabus, id_event FROM mahasiswa WHERE id =?
`

func (q *Queries) GetMahasiswa(ctx context.Context, id int32) (Mahasiswa, error) {
	row := q.db.QueryRowContext(ctx, getMahasiswa, id)
	var i Mahasiswa
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Nomorhp,
		&i.Email,
		&i.UrlFoto,
		&i.NamaLengkap,
		&i.TanggalLahir,
		&i.JenisKelamin,
		&i.AsalKampus,
		&i.Nim,
		&i.Jurusan,
		&i.TahunMasuk,
		&i.KotaKabupaten,
		&i.IDTantangan,
		&i.IDSilabus,
		&i.IDEvent,
	)
	return i, err
}

const getMahasiswaByAuth = `-- name: GetMahasiswaByAuth :one
SELECT id, username, password, nomorhp, email, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, asal_kampus, nim, jurusan, tahun_masuk, kota_kabupaten, id_tantangan, id_silabus, id_event FROM mahasiswa WHERE username = ? AND password = ? OR email = ? AND password = ?
`

type GetMahasiswaByAuthParams struct {
	Username   sql.NullString `json:"username"`
	Password   sql.NullString `json:"password"`
	Email      sql.NullInt32  `json:"email"`
	Password_2 sql.NullString `json:"password_2"`
}

func (q *Queries) GetMahasiswaByAuth(ctx context.Context, arg GetMahasiswaByAuthParams) (Mahasiswa, error) {
	row := q.db.QueryRowContext(ctx, getMahasiswaByAuth,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Password_2,
	)
	var i Mahasiswa
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Nomorhp,
		&i.Email,
		&i.UrlFoto,
		&i.NamaLengkap,
		&i.TanggalLahir,
		&i.JenisKelamin,
		&i.AsalKampus,
		&i.Nim,
		&i.Jurusan,
		&i.TahunMasuk,
		&i.KotaKabupaten,
		&i.IDTantangan,
		&i.IDSilabus,
		&i.IDEvent,
	)
	return i, err
}

const updateMahasiswa = `-- name: UpdateMahasiswa :exec
UPDATE mahasiswa SET email = ?, url_foto = ?, nama_lengkap = ?, tanggal_lahir = ?, jenis_kelamin = ?, asal_kampus = ? , nim = ?, jurusan = ?, tahun_masuk = ?, kota_kabupaten = ? WHERE id = ?
`

type UpdateMahasiswaParams struct {
	Email         sql.NullInt32  `json:"email"`
	UrlFoto       sql.NullString `json:"url_foto"`
	NamaLengkap   sql.NullString `json:"nama_lengkap"`
	TanggalLahir  sql.NullString `json:"tanggal_lahir"`
	JenisKelamin  sql.NullString `json:"jenis_kelamin"`
	AsalKampus    sql.NullString `json:"asal_kampus"`
	Nim           sql.NullString `json:"nim"`
	Jurusan       sql.NullString `json:"jurusan"`
	TahunMasuk    sql.NullString `json:"tahun_masuk"`
	KotaKabupaten sql.NullString `json:"kota_kabupaten"`
	ID            int32          `json:"id"`
}

func (q *Queries) UpdateMahasiswa(ctx context.Context, arg UpdateMahasiswaParams) error {
	_, err := q.db.ExecContext(ctx, updateMahasiswa,
		arg.Email,
		arg.UrlFoto,
		arg.NamaLengkap,
		arg.TanggalLahir,
		arg.JenisKelamin,
		arg.AsalKampus,
		arg.Nim,
		arg.Jurusan,
		arg.TahunMasuk,
		arg.KotaKabupaten,
		arg.ID,
	)
	return err
}

const viewMahasiswa = `-- name: ViewMahasiswa :many
SELECT id, username, password, nomorhp, email, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, asal_kampus, nim, jurusan, tahun_masuk, kota_kabupaten, id_tantangan, id_silabus, id_event FROM mahasiswa
`

func (q *Queries) ViewMahasiswa(ctx context.Context) ([]Mahasiswa, error) {
	rows, err := q.db.QueryContext(ctx, viewMahasiswa)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mahasiswa
	for rows.Next() {
		var i Mahasiswa
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Nomorhp,
			&i.Email,
			&i.UrlFoto,
			&i.NamaLengkap,
			&i.TanggalLahir,
			&i.JenisKelamin,
			&i.AsalKampus,
			&i.Nim,
			&i.Jurusan,
			&i.TahunMasuk,
			&i.KotaKabupaten,
			&i.IDTantangan,
			&i.IDSilabus,
			&i.IDEvent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
