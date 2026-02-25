CREATE TABLE psql_service (
    id uuid PRIMARY KEY,
    project_id uuid NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    service_id TEXT NOT NULL,
    name TEXT NOT NULL,
    app_name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    db_name TEXT NOT NULL,
    db_user TEXT NOT NULL,
    db_password TEXT NOT NULL,
    image TEXT NOT NULL,
    internal_url TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);