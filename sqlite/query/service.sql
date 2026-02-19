-- name: GetAllServices :many
SELECT p.id,p.name
FROM project p
JOIN service s ON p.id = s.project_id
WHERE p.id = ?;

-- name: CheckProjectHasServices :one
SELECT CAST(EXISTS (
    SELECT 1 FROM service s
    JOIN project p ON s.project_id = p.id
    WHERE p.id = ?
) AS BOOLEAN);

-- name: GetService :one
SELECT *
FROM service
WHERE id = ?;

-- name: CreateService :one
INSERT INTO service (name,project_id)
VALUES (?,?)
RETURNING id;

-- name: DeleteService :exec
DELETE FROM service
WHERE id = ?;