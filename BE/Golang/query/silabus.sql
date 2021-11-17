-- name: ListSilabus :many
SELECT * FROM silabus;

-- name: UpdateSilabus :exec
UPDATE silabus SET task_silabus = ?, ujian_silabus = ?, skor_silabus = ? WHERE id = ?;