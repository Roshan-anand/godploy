-- name: InsertOrg :one
INSERT INTO organization (name)
VALUES (?)
RETURNING id;

