package models

import (
	"time"
)

type Event struct {
	ID             int       `gorm:"primary_key"json:"id"`
	JudulEvent     string    `json:"judul_event"`
	DeskripsiEvent string    `json:"deskripsi_event"`
	KriteriaEvent  string    `json:"kriteria_event"`
	TanggalEvent   time.Time `json:"tanggal_event"`
	IDAdmin        int       `json:"id_admin"`
}
