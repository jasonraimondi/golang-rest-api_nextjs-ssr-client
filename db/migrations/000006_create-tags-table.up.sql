CREATE TABLE tags
(
    id   BIGINT       NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE photo_tag
(
    photo_id UUID   NOT NULL,
    tag_id   BIGINT NOT NULL,
    PRIMARY KEY (photo_id, tag_id),
    FOREIGN KEY (photo_id) REFERENCES photos (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON UPDATE CASCADE ON DELETE CASCADE
);
