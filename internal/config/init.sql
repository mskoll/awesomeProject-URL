CREATE TABLE url
(
    id SERIAL PRIMARY KEY ,
    url VARCHAR(255) NOT NULL ,
    short_url VARCHAR(50)
);