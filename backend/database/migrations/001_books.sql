create extension if not exists "uuid-ossp";
create table books (
    uuid uuid primary key default uuid_generate_v4(),

    title text not null,
    author text not null,
    isbn text unique not null,
    genre text not null,
    publisher text not null,

    year numeric,
    language text,
    cover_type text,
    color boolean,
    page_count numeric,
    description text
);

---- create above / drop below ----

drop table books;
