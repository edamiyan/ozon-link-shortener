CREATE TABLE IF NOT EXISTS links
(
    id      serial        not null,
    baseURL varchar(1024) not null,
    token   varchar(10)   not null
);