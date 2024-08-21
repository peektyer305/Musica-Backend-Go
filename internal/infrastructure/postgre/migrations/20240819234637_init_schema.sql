-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE DATABASE app;
create table "app"."users"
(
	id UUID PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_icon_url VARCHAR(255),
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

create table "app"."posts"
(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(50) NOT NULL,
    content TEXT,
    music_url VARCHAR(255) NOT NULL,
    image_url VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table "app"."users";
-- +goose StatementEnd