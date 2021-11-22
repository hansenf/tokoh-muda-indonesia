-- name: CheckAdmin :one
SELECT * FROM admin WHERE role = "admin";

-- name: CreateAdmin :exec
INSERT INTO admin (
    username, password, role
)VALUES(
    ?,?,?
);