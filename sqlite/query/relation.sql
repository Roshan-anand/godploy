-- name: LinkUserNOrg :exec
INSERT INTO user_organization (user_id, organization_id)
VALUES (?, ?);
