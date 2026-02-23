-- name: CreateOrg :one
INSERT INTO organization (id, name)
VALUES (?, ?)
RETURNING id;

-- name: DeleteOrg :exec
DELETE FROM organization
WHERE id = ?;

-- name: GetAllOrg :many
SELECT o.id,o.name
FROM organization o
JOIN user_organization uo ON o.id = uo.organization_id
WHERE uo.user_email = ?;

-- name: CheckProjectExist :one
SELECT CAST(EXISTS(
    SELECT 1 FROM project p
    JOIN organization o ON o.id = p.organization_id
    WHERE o.id = @org_id  AND p.name = @project_name
) AS BOOLEAN );

-- name: GetAllProjects :many
SELECT p.id,p.name
FROM organization o
JOIN project p ON o.id = p.organization_id
WHERE o.id = @org_id;

-- name: CreateProject :one
INSERT INTO project (id,name,organization_id)
VALUES (?,?,@org_id)
RETURNING id,name;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = ?;

-- name: CheckProjectHasServices :one
SELECT CAST(EXISTS (
    SELECT 1 FROM project p
    JOIN psql_service psql ON psql.project_id = p.id
    WHERE p.id = ?
) AS BOOLEAN);