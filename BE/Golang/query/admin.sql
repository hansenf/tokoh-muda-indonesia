-- name: CreateAdmin :exec
INSERT INTO admin(
    username, password, role
) VALUES(
    ?,?,?
);

-- name: GetAdmin :one
SELECT * FROM admin WHERE username = ? AND password = ?;

-- name: GetOneAdmin :one
SELECT * FROM admin WHERE username = "admin";