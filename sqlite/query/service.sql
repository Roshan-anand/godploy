-- name: CreatePsqlService :one
INSERT INTO psql_service (psql_id,project_id, name, app_name, description, db_name, db_user, db_password, image, internal_url)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetPsqlServiceById :one
SELECT *
FROM psql_service
WHERE psql_id = ?;

-- name: SetPsqlServiceId :exec
UPDATE psql_service
SET serviceid = ?
WHERE psql_id = ?;

-- name: DeletePsqlService :exec
DELETE FROM psql_service
WHERE psql_id = ?