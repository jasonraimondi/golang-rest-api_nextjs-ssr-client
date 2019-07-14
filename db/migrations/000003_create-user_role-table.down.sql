ALTER TABLE user_role DROP CONSTRAINT user_role_user_id_fkey;
ALTER TABLE user_role DROP CONSTRAINT user_role_role_id_fkey;
DROP TABLE IF EXISTS user_role;
