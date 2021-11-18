package models

import (
	"time"
)

type Mahasiswa struct {
	ID            int       `gorm:"primary_key" json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Nomorhp       string    `json:"nomorhp"`
	Email         string    `json:"email"`
	UrlFoto       string    `json:"url_foto"`
	NamaLengkap   string    `json:"nama_lengkap"`
	TanggalLahir  time.Time `json:"tanggal_lahir"`
	JenisKelamin  string    `json:"jenis_kelamin"`
	AsalKampus    string    `json:"asal_kampus"`
	Nim           int       `json:"nim"`
	Jurusan       string    `json:"jurusan"`
	TahunMasuk    int       `json:"tahun_masuk"`
	KotaKabupaten string    `json:"kota_kabupaten"`
	IDTantangan   int       `json:"id_tantangan"`
	IDSilabus     int       `json:"id_silabus"`
	IDEvent       int       `json:"id_event"`
}
