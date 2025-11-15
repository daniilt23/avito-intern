-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS teams (
    team_name VARCHAR(100) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(20) PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    team_name VARCHAR(100) NOT NULL REFERENCES teams(team_name) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_users_team_active ON users(team_name, is_active);

CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_requests (
    pull_request_id VARCHAR(20) PRIMARY KEY,
    pull_request_name VARCHAR(255) NOT NULL,
    status pr_status NOT NULL DEFAULT 'OPEN',
    author_id VARCHAR(20) NOT NULL REFERENCES users(user_id),
    need_more_reviewers BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP(0) DEFAULT NOW(),
    merged_at TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS idx_pull_requests_status ON pull_requests(status);

CREATE TABLE IF NOT EXISTS pr_users (
    id SERIAL PRIMARY KEY,
    pull_request_id VARCHAR(20) NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    user_id VARCHAR(20) NOT NULL REFERENCES users(user_id)
);

CREATE INDEX IF NOT EXISTS idx_pr_users_pull_request_id ON pr_users(pull_request_id);
CREATE INDEX IF NOT EXISTS idx_pr_users_user_id ON pr_users(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
