CREATE TABLE redirect_session (
    state TEXT PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES user(id) ON DELETE CASCADE,
    org_id uuid NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    expires_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE github_app (
    id uuid PRIMARY KEY,
    name TEXT NOT NULL,
    organization_id uuid NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    app_id INTEGER NOT NULL,
    installation_id INTEGER,
    pem_key TEXT NOT NULL,
    webhook_secret TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);