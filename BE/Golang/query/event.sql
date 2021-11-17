-- name: CreateEvent :exec
INSERT INTO event(
    judul_event, deskripsi_event, kriteria_event, tanggal_event
) VALUES(
    ?,?,?,?
);

-- name: UpdateEvent :exec
UPDATE event SET judul_event=?, deskripsi_event=?, kriteria_event=?, tanggal_event=? WHERE id = ?;;

-- name: DeleteEvent :exec
DELETE FROM event WHERE id = ?;

-- name: GetEvent :many
SELECT * FROM event;
