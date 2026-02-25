-- name: CreatePsqlService :one
INSERT INTO psql_service (id,project_id,service_id, name, app_name, description, db_name, db_user, db_password, image, internal_url)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetPsqlServiceById :one
SELECT *
FROM psql_service
WHERE id = ?;

-- name: SetPsqlServiceId :exec
UPDATE psql_service
SET service_id = ?
WHERE id = ?;

-- name: DeletePsqlService :exec
DELETE FROM psql_service
WHERE id = ?