-- name: CreateUser :one
INSERT INTO user (name, email, hash_pass, role)
VALUES (?, ?, ?, ?)
RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM user
WHERE email = ?;

-- name: RemoveUser :exec
DELETE FROM user
WHERE id = ?;

-- name: AdminExists :one
SELECT CAST(EXISTS (
    SELECT 1 FROM user
    WHERE role = 'admin'
) AS BOOLEAN);
