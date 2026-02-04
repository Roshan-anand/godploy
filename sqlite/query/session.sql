-- name: CreateSession :exec
INSERT INTO session (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetSessionByToken :one
SELECT u.id,u.email,u.name,s.expires_at,s.created_at
FROM session s
INNER JOIN user u ON s.user_id = u.id
WHERE s.token = ?;

-- name: RemoveSessionByUID :exec
DELETE FROM session
WHERE user_id = ?;