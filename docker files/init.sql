CREATE TABLE accounts (
    ID SERIAL PRIMARY KEY,
    email varchar(320) UNIQUE,
    password  bytea
);

--CREATE INDEX accounts_email_index on accounts using hash(email);