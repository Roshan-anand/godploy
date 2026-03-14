-- name: CreateRedirectSession :exec
INSERT INTO redirect_session (state, user_id, org_id, expires_at)
VALUES (?, ?, ?, ?);

-- name: GetRedirectSession :one
SELECT state, user_id, org_id, expires_at, created_at
FROM redirect_session
WHERE state = ?;

-- name: DeleteRedirectSession :exec
DELETE FROM redirect_session
WHERE state = ?;

-- name: CreateGithubApp :exec
INSERT INTO github_app (id, organization_id,app_id, pem_key, webhook_secret)
VALUES (?, ?, ?, ?, ?);

-- name: GetGithubApp :one
SELECT * FROM github_app
WHERE organization_id = ?;

-- name: InsertInstallationID :exec
UPDATE github_app
SET installation_id = ?, updated_at = CURRENT_TIMESTAMP
WHERE organization_id = ?;
