-- name: CheckProjectExist :one
SELECT CAST(EXISTS(
    SELECT 1 FROM project p
    JOIN organization o ON o.id = p.organization_id
    WHERE o.id = @org_id  AND p.name = @project_name
) AS BOOLEAN );

-- name: GetAllProjects :many
SELECT p.id,p.name,p.description
FROM organization o
JOIN project p ON o.id = p.organization_id
WHERE o.id = @org_id;

-- name: CreateProject :one
INSERT INTO project (id,name,description,organization_id)
VALUES (?,?,?,@org_id)
RETURNING id,name,description;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = ?;

-- name: CheckProjectHasServices :one
SELECT CAST(EXISTS(
    SELECT 1
    FROM psql_service ps
    WHERE ps.project_id = @project_id
    UNION
    SELECT 1
    FROM app_service aps
    WHERE aps.project_id = @project_id
) AS BOOLEAN);
