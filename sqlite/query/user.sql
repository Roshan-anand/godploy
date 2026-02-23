-- name: CreateUser :one
INSERT INTO user (id,name, email, hash_pass, role)
VALUES (?,?, ?, ?, ?)
RETURNING id;

-- name: GetUserByEmail :one
SELECT
   u.id,
   u.name,
   u.email,
   u.hash_pass,
   u.role,
   CAST(
    json_group_array(
        json_object('id', o.id, 'name', o.name, 'created_at',  strftime('%Y-%m-%dT%H:%M:%SZ', o.created_at))
    ) AS TEXT
   ) AS orgs
FROM user u
JOIN user_organization uo ON u.email = uo.user_email
JOIN organization o ON uo.organization_id = o.id
WHERE u.email = ?
GROUP BY u.name, u.email, u.hash_pass,u.role;

-- name: RemoveUser :exec
DELETE FROM user
WHERE id = ?;

-- name: AdminExists :one
SELECT CAST(EXISTS (
    SELECT 1 FROM user
    WHERE role = 'admin'
) AS BOOLEAN);
