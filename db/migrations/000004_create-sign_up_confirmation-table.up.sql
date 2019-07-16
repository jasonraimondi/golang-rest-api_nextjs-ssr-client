CREATE TABLE sign_up_confirmation
(
    token      UUID   NOT NULL,
    user_id    UUID   NOT NULL,
    created_at BIGINT NOT NULL,
    PRIMARY KEY (token),
    FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
