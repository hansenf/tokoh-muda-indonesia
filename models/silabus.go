package models

type CreateSilabusParams struct {
	ID              int    `gorm:"primary_key"json:"id"`
	JudulSilabus    string `json:"judul_silabus"`
	Definisi        string `json:"definisi"`
	FungsiSilabus   string `json:"fungsi_silabus"`
	Deskripsi       string `json:"deskripsi"`
	UrlVideoSilabus string `json:"url_video_silabus"`
	TaskSilabus     string `json:"task_silabus"`
	UjianSilabus    string `json:"ujian_silabus"`
	SkorSilabus     int    `json:"skor_silabus"`
	IDAdmin         int    `json:"id_admin"`
}
