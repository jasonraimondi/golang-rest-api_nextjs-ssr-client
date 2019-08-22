ALTER TABLE photo_app
    DROP CONSTRAINT photo_app_photo_id_fkey;
ALTER TABLE photo_app
    DROP CONSTRAINT photo_app_app_id_fkey;
DROP TABLE IF EXISTS photo_app;

DROP TABLE IF EXISTS apps;
