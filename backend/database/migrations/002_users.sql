create table users (
    id serial primary key,

    email text unique not null,
    login text unique not null,
    name text not null,

    salt bytea not null,
    hash bytea not null
);

---- create above / drop below ----

drop table users;
