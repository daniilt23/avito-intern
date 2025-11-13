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

CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE IF NOT EXISTS pull_requests (
    pull_request_id VARCHAR(20) PRIMARY KEY,
    pull_request_name VARCHAR(255) NOT NULL,
    status pr_status NOT NULL DEFAULT 'OPEN',
    author_id VARCHAR(20) NOT NULL REFERENCES users(user_id),
    need_more_reviewers BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NULL,
    merged_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS pr_users (
    id SERIAL PRIMARY KEY,
    pull_request_id VARCHAR(20) NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    user_id VARCHAR(20) NOT NULL REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
