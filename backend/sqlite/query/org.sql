-- name: GetCurrentOrg :one
SELECT o.id,o.name 
FROM user u
JOIN organization o ON o.id = u.current_org_id
WHERE u.email = ?;
