-- +goose Up
-- +goose StatementBegin

-- users テーブル作成
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    username varchar(50) UNIQUE NOT NULL,
    user_info text,
    user_icon_url varchar(255),
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

-- user_privates テーブル作成
CREATE TABLE IF NOT EXISTS user_privates (
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

-- user_privates テーブル削除
DROP TABLE IF EXISTS user_privates;

-- users テーブル削除
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
