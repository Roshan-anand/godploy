-- name: CheckProjectExist :one
SELECT CAST(EXISTS(
    SELECT 1
    FROM project p
    JOIN user u ON u.current_org_id = p.organization_id
    WHERE u.email = ? AND p.name = @project_name
) AS BOOLEAN );

-- name: GetAllProjects :many
SELECT p.id,p.name,p.description
FROM project p
JOIN user u ON u.current_org_id = p.organization_id
WHERE u.email = ?;

-- name: CreateProject :one
INSERT INTO project (id,name,description,organization_id)
SELECT ?, ?, ?, u.current_org_id
FROM user u
WHERE u.email = ?
LIMIT 1
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
