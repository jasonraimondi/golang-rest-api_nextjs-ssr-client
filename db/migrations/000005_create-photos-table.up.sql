CREATE TABLE photos
(
    id            UUID         NOT NULL,
    original_name VARCHAR(100) NOT NULL,
    content_type  varchar(100) NOT NULL,
    file_size     BIGINT       NOT NULL,
    user_id       UUID         NOT NULL,
    created_at    BIGINT       NOT NULL,
    modified_at   BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
