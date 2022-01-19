SET timezone ='+3';

/* profile table */

CREATE TABLE if not exists Profile
(
    user_id integer NOT NULL PRIMARY KEY,
    user_mail VARCHAR(32),
    username VARCHAR(32),
    user_surname VARCHAR(64),
    user_role VARCHAR(64)
);