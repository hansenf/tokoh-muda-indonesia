-- name: CreateUser :exec
INSERT INTO user (
  username,email, nomorhp, password
) VALUES (
  ?, ?,?,?
);

-- name: GetUserByID :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserByEmailAndPassword :one
SELECT * FROM user WHERE email = ? AND password = ?;