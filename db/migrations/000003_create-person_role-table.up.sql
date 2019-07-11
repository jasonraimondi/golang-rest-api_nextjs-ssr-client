CREATE TABLE person_role
(
    user_id integer NOT NULL,
    role_id integer NOT NULL,
    grant_date timestamp without time zone,
    PRIMARY KEY (user_id, role_id)
--     CONSTRAINT person_role_role_id_fkey FOREIGN KEY (role_id)
--         REFERENCES roles (id) MATCH SIMPLE
--         ON UPDATE NO ACTION ON DELETE NO ACTION,
--     CONSTRAINT person_role_person_id_fkey FOREIGN KEY (user_id)
--         REFERENCES persons (id) MATCH SIMPLE
--         ON UPDATE NO ACTION ON DELETE NO ACTION
);
