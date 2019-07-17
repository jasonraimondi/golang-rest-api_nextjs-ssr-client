CREATE TABLE photos
(
    id           UUID         NOT NULL,
    sha256       VARCHAR(64)  NOT NULL,
    file_name    VARCHAR(100) NOT NULL,
    relative_url VARCHAR(255) NOT NULL,
    mime_type    varchar(100) NOT NULL,
    file_size    BIGINT       NOT NULL,
    width        INT,
    height       INT,
    user_id      UUID         NOT NULL,
    created_at   BIGINT       NOT NULL,
    modified_at  BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
