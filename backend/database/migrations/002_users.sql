create table users (
    id serial primary key,

    email text unique not null,
    login text unique not null,
    name text not null,

    salt varchar(16) not null,
    hash varchar(32) not null
);

---- create above / drop below ----

drop table users;
