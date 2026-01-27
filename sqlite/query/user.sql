-- name: AddUser :exec
INSERT INTO user (name, email, hash_pass, role)
VALUES (?, ?, ?, ?);

-- name: GetUserByEmail :one
SELECT * FROM user
WHERE email = ?;

-- name: RemoveUser :exec
DELETE FROM user
WHERE id = ?;
