CREATE TABLE users 
(
    id            serial       not null unique,
    email         varchar(255) not null,
    username      varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE files 
(
    id           serial       not null unique,
    size         int          not null,
    url          varchar(255) not null,
    created_date timestamp    not null,
    owner_id     int          references users(id) on delete cascade not null
);