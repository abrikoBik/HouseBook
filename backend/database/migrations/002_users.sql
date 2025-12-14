create table users (
    id serial primary key,
    login text unique not null,
    name text not null
);

---- create above / drop below ----

drop table users;
