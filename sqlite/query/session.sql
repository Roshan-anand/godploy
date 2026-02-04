-- name: CreateSession :exec
INSERT INTO session (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetSessionByToken :one
SELECT u.email,s.expires_at,s.created_at
FROM session s
INNER JOIN user u ON s.user_id = u.id
WHERE s.token = ?;
