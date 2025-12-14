create extension if not exists "uuid-ossp";
create table books (
    uuid uuid primary key default uuid_generate_v4(),
    title text not null,
    author text not null,
    isbn text unique not null
);

---- create above / drop below ----

drop table books;
