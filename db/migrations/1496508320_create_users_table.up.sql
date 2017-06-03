CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    firstName varchar(40) NOT NULL,
    email varchar(40) NOT NULL,
    password varchar(40) NOT NULL
);
