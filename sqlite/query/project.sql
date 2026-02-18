-- name: CreateOrg :one
INSERT INTO organization (name)
VALUES (?)
RETURNING id;

-- name: DeleteOrg :exec
DELETE FROM organization
WHERE id = ?;

-- name: GetAllOrg :many
SELECT o.id,o.name
FROM organization o
JOIN user_organization uo ON o.id = uo.organization_id
WHERE uo.user_email = ?;

-- name: GetAllProjects :many
SELECT p.id,p.name
FROM organization o
JOIN project p ON o.id = p.organization_id
WHERE o.id = ?;

-- name: CreateProject :one
INSERT INTO project (name,organization_id)
VALUES (?,?)
RETURNING id,name;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = ?;

-- name: GetAllServices :many
SELECT p.id,p.name
FROM project p
JOIN service s ON p.id = s.project_id
WHERE p.id = ?;

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
