-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- スキーマが存在しない場合は作成

-- usersテーブルを作成
CREATE TABLE IF NOT EXISTS "users"
(
    id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_icon_url VARCHAR(255),
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- postsテーブルを作成
CREATE TABLE IF NOT EXISTS "posts"
(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(50) NOT NULL,
    content TEXT,
    music_url VARCHAR(255) NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "users"(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

-- postsテーブルを削除
DROP TABLE IF EXISTS "posts";

-- usersテーブルを削除
DROP TABLE IF EXISTS "users";

-- スキーマを削除（スキーマごと削除する場合はコメントを解除）
-- DROP SCHEMA IF EXISTS "app" CASCADE;
-- +goose StatementEnd
