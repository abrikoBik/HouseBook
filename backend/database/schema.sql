create extension if not exists "uuid-ossp";

-- TODO: original owner + current owner
create table books (
    id uuid primary key default uuid_generate_v4(),

    title text not null,
    author text not null,
    isbn text unique not null,
    genre text not null,
    publisher text not null,

    year numeric,
    language text,
    cover_type text,
    page_count numeric,
    description text
);

create table users (
    id bigserial primary key,

    email text unique not null,
    login text unique not null,
    name text not null,

    salt bytea not null,
    hash bytea not null
);

create table trades (
    id uuid primary key default uuid_generate_v4(),

    owner_id bigserial not null references users(id),
    buyer_id bigserial not null references users(id),

    created_at timestamp not null default current_timestamp,
    completed_at timestamp
);

create table trade_books (
    trade_id uuid references trades(id) on delete cascade,
    book_id uuid references books(id),
    owner_id bigserial not null references users(id),
    primary key (trade_id, book_id)
);
