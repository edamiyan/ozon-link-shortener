CREATE TABLE IF NOT EXISTS links
(
    id      serial        not null,
    base_url varchar(1024) not null,
    token   varchar(10)   not null,
    UNIQUE (token)
);