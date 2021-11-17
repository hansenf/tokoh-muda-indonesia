-- name: CreateEvent :exec
INSERT INTO event(
    judul_event, deskripsi_event, kriteria_event, tanggal_event
) VALUES(
    ?,?,?,?
);

-- name: DeleteEvent :exec
DELETE FROM event WHERE id = ?;

-- name: GetEvent :many
SELECT * FROM event;
