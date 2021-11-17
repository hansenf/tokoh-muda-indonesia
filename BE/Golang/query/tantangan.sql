-- name: ListTantangan :many
SELECT * FROM tantangan;

-- name: UpdateTantangan :exec
UPDATE tantangan SET task_tantangan = ?, ujian_tantangan = ?, skor_tantangan = ? WHERE id = ?;