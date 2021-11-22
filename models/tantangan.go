package models

type Tantangan struct {
	ID                int    `gorm:"primary_key"json:"id"`
	JudulTantangan    string `json:"judul_tantangan"`
	Tema              string `json:"tema"`
	Latar             string `json:"latar"`
	UrlVideoTantangan string `json:"url_video_tantangan"`
	TaskTantangan     string `json:"task_tantangan"`
	UjianTantangan    string `json:"ujian_tantangan"`
	SkorTantangan     int    `json:"skor_tantangan"`
	IDAdmin           int    `json:"id_admin"`
}
