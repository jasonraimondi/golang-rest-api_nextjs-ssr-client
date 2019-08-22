CREATE TABLE apps
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE photo_app
(
    photo_id    UUID   NOT NULL,
    app_id BIGINT NOT NULL,
    PRIMARY KEY (photo_id, app_id),
    CONSTRAINT unique_photo_app UNIQUE (photo_id, app_id),
    FOREIGN KEY (photo_id) REFERENCES photos (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (app_id) REFERENCES apps (id) ON UPDATE CASCADE ON DELETE CASCADE
);
