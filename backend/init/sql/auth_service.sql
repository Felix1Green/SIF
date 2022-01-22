SET timezone ='+3';

/* users table */

CREATE TABLE if not exists Users
(
    user_id serial NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL
);