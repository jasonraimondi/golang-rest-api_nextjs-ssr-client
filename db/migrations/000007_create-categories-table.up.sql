CREATE TABLE categories
(
    id   BIGINT       NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE photo_category
(
    photo_id    UUID   NOT NULL,
    category_id BIGINT NOT NULL,
    PRIMARY KEY (photo_id, category_id),
    CONSTRAINT unique_photo_category UNIQUE (photo_id, category_id),
    FOREIGN KEY (photo_id) REFERENCES photos (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON UPDATE CASCADE ON DELETE CASCADE
);
