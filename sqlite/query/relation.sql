-- name: LinkUserNOrg :exec
INSERT INTO user_organization (user_email, organization_id)
VALUES (?, ?);

-- name: CheckUserOrgExists :one
SELECT CAST(EXISTS(
    SELECT 1 FROM user_organization uo
    WHERE uo.user_email = ? AND uo.organization_id = ?
)AS BOOLEAN)
