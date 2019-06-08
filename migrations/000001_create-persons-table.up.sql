CREATE TABLE persons
(
    id            VARCHAR(32) PRIMARY KEY,
    first_name    VARCHAR(255),
    last_name     VARCHAR(255),
    password_hash VARCHAR(255),
    email         VARCHAR(355) UNIQUE NOT NULL,
    created_at    BIGINT NOT NULL,
    modified_at   BIGINT
);
