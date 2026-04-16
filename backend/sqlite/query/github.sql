-- name: CreateRedirectSession :exec
INSERT INTO redirect_session (state, user_id, org_id, expires_at)
VALUES  (?, ?, ?, ?);

-- name: GetRedirectSession :one
SELECT state, user_id, org_id, gh_app_id, expires_at, created_at
FROM redirect_session
WHERE state = ?;

-- name: GetRedirectSessionGhAppID :one
SELECT gh_app_id
FROM redirect_session
WHERE state = ?;

-- name: UpdateRedirectSession :exec
UPDATE redirect_session
SET gh_app_id = ?
WHERE state = ?;

-- name: DeleteRedirectSession :exec
DELETE FROM redirect_session
WHERE state = ?;

-- name: CreateGithubApp :one
INSERT INTO github_app (id, name, organization_id,app_id, pem_key, webhook_secret)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING app_id;

-- name: GetGhAppByAppId :one
SELECT * FROM github_app
WHERE app_id = ?;

-- name: GetAllGhAppsByEmail :many
SELECT gh.name, gh.app_id, gh.created_at
FROM user u
JOIN github_app gh ON u.current_org_id = gh.organization_id
WHERE u.email = ?;

-- name: InsertInstallationID :exec
UPDATE github_app
SET installation_id = ?, updated_at = CURRENT_TIMESTAMP
WHERE app_id = ?;

-- name: DeleteGithubApp :exec
DELETE FROM github_app
WHERE app_id = ?;
