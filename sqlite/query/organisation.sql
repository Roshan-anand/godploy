-- name: CreateOrg :one
INSERT INTO organization (name)
VALUES (?)
RETURNING id;

