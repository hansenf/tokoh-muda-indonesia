-- name: CreateSilabus :exec
INSERT INTO silabus(
    id, judul_silabus, definisi, fungsi_silabus, deskripsi, url_video_silabus,task_silabus, ujian_silabus,id_admin
)VALUES(
    ?,?,?,?,?,?,?,?,?
);