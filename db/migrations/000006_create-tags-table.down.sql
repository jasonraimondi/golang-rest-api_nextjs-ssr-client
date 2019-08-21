ALTER TABLE photo_tag
    DROP CONSTRAINT photo_tag_photo_id_fkey;
ALTER TABLE photo_tag
    DROP CONSTRAINT photo_tag_tag_id_fkey;
DROP TABLE IF EXISTS photo_tag;

DROP TABLE IF EXISTS tags;
