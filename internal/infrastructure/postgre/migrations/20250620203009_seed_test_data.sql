-- +goose Up
-- +goose StatementBegin

INSERT INTO "users" (id, username, user_info, user_icon_url, user_client_id, created_at, updated_at) VALUES
('01978d0d-e0cf-78d6-bf40-95a1e37d72e4', 'Achro', 'Developer of musica.', 'https://tmicfifqruiiqpqiemji.supabase.co/storage/v1/object/public/usericon//images.jpg', 'Developer', NOW(), NOW()),
('01978d0d-e0cf-7b51-99e4-851826df2392', 'BisRa', 'SubAccount of Achro', 'https://tmicfifqruiiqpqiemji.supabase.co/storage/v1/object/public/usericon//2222222.jpg', 'SubAccount', NOW(), NOW());

INSERT INTO "user_privates" (id, user_id, mail_address, password, token, created_at, updated_at) VALUES
('01978d0d-e0cf-7f7e-88ca-bbf0d7f6ac41', '01978d0d-e0cf-78d6-bf40-95a1e37d72e4', 'pektyer305@gmail.com', 'password1', '1111111111111', NOW(), NOW()),
('01978d0d-e0cf-7fc4-9b78-1a44f405a612', '01978d0d-e0cf-7b51-99e4-851826df2392', 'bisra@example.com', 'password2', '11111111111111111', NOW(), NOW());

INSERT INTO "posts" (id, user_id, title, content, music_url, image_url, created_at, updated_at) VALUES
('01978d17-4d61-7a7d-b28c-17e5280c6f21', '01978d0d-e0cf-78d6-bf40-95a1e37d72e4', 'first', 'This is first post.', 'https://youtu.be/s4Afw3_jc_k?si=n3qvKpDSsjeeCld30', NULL, NOW(), NOW());

-- +goose StatementEnd