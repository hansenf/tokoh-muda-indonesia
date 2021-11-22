-- name: CreateTantangan :exec
INSERT INTO tantangan(
    id, judul_tantangan, tema, latar, url_video_tantangan, task_tantangan, ujian_tantangan, id_admin
)VALUES(
    ?,?,?,?,?,?,?,?
);