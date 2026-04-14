ALTER TABLE app_service ADD COLUMN git_provider TEXT NOT NULL DEFAULT '';
ALTER TABLE app_service ADD COLUMN git_repo_id TEXT NOT NULL DEFAULT '';
ALTER TABLE app_service ADD COLUMN git_repo_name TEXT NOT NULL DEFAULT '';
