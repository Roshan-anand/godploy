-- name: LinkUserNOrg :exec
INSERT INTO user_organization (user_email, organization_id)
VALUES (?, ?);
