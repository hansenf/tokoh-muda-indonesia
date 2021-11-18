package models

type Leaderboard struct {
	Ranking     int `json:"gorm:"primary_key"ranking"`
	IDMahasiswa int `json:"id_mahasiswa"`
	IDTantangan int `json:"id_tantangan"`
	IDSilabus   int `json:"id_silabus"`
}
