-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- 2) users テーブルを作成
CREATE TABLE IF NOT EXISTS "users" (
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_info TEXT,
    user_icon_url VARCHAR(255),
    user_client_id TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- user_privates テーブル作成
CREATE TABLE IF NOT EXISTS "user_privates" (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL UNIQUE,
    mail_address varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    token varchar(255),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp NULL,
    CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	  REFERENCES users(id)
      ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_privates";
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
