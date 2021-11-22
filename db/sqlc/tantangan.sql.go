// Code generated by sqlc. DO NOT EDIT.
// source: tantangan.sql

package db

import (
	"context"
)

const createTantangan = `-- name: CreateTantangan :exec
INSERT INTO tantangan(
    id, judul_tantangan, tema, latar, url_video_tantangan, task_tantangan, ujian_tantangan, id_admin
)VALUES(
    ?,?,?,?,?,?,?,?
)
`

type CreateTantanganParams struct {
	ID                int32  `json:"id"`
	JudulTantangan    string `json:"judul_tantangan"`
	Tema              string `json:"tema"`
	Latar             string `json:"latar"`
	UrlVideoTantangan string `json:"url_video_tantangan"`
	TaskTantangan     string `json:"task_tantangan"`
	UjianTantangan    string `json:"ujian_tantangan"`
	IDAdmin           int    `json:"id_admin"`
}

func (q *Queries) CreateTantangan(ctx context.Context, arg CreateTantanganParams) error {
	_, err := q.db.ExecContext(ctx, createTantangan,
		arg.ID,
		arg.JudulTantangan,
		arg.Tema,
		arg.Latar,
		arg.UrlVideoTantangan,
		arg.TaskTantangan,
		arg.UjianTantangan,
		arg.IDAdmin,
	)
	return err
}
