CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    first    VARCHAR(255),
    last     VARCHAR(255),
    password_hash VARCHAR(255),
    email         VARCHAR(255) UNIQUE NOT NULL,
    is_verified   BOOLEAN DEFAULT FALSE,
    created_at    BIGINT NOT NULL,
    modified_at   BIGINT
);
