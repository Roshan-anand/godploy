-- name: CreateDeployment :one
INSERT INTO deployments (id, service_id, name, status)
VALUES (?, ?, ?, ?)
RETURNING id;

-- name: GetDeploymentByID :one
SELECT id, service_id, name, status, created_at
FROM deployments
WHERE id = ?;

-- name: GetDeploymentsByServiceID :many
SELECT id, service_id, name, status, created_at
FROM deployments
WHERE service_id = ?
ORDER BY created_at DESC;

-- name: GetAllDeploymentIdsByServiceID :many
SELECT id
FROM deployments
WHERE service_id = ?;

-- name: GetDeploymentStatus :one
SELECT status
FROM deployments
WHERE id = ?;

-- name: UpdateDeploymentStatus :exec
UPDATE deployments
SET status = ?
WHERE id = ?;

-- name: DeleteDeploymentByID :exec
DELETE FROM deployments
WHERE id = ?;
