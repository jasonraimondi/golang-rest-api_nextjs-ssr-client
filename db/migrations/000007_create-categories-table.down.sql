ALTER TABLE photo_category
    DROP CONSTRAINT photo_category_photo_id_fkey;
ALTER TABLE photo_category
    DROP CONSTRAINT photo_category_category_id_fkey;
DROP TABLE IF EXISTS photo_category;

DROP TABLE IF EXISTS categories;
