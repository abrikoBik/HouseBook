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
    description text,

    image_uuid uuid,
    link text
);

create table users (
    id uuid primary key default uuidv7(),

    email text unique not null,
    login text unique not null,
    name text not null,

    salt bytea not null,
    hash bytea not null,

    image_uuid uuid,
    rating real not null default 5 check (rating >= 1 and rating <= 5),

    now_reading_id uuid references libraries(id),
    finished_reading_id uuid references libraries(id),
    trade_ready_id uuid references libraries(id)
);

create table trades (
    id uuid primary key default uuid_generate_v4(),

    owner_id uuid not null references users(id),
    buyer_id uuid not null references users(id),
    owner_confirmed boolean not null default false,
    buyer_confirmed boolean not null default false,

    created_at timestamp not null default current_timestamp,
    completed_at timestamp
);

create table trade_books (
    trade_id uuid references trades(id) on delete cascade,
    book_id uuid references books(id),
    owner_id uuid not null references users(id),
    primary key (trade_id, book_id)
);

create table libraries (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id),
    name text not null default 'Библиотека'
);

create table library_books (
    library_id uuid references libraries(id) on delete cascade,
    book_id uuid references books(id),
    quantity numeric not null default 1,
    primary key (library_id, book_id)
);

create table comments (
    id uuid default uuid_generate_v4(),
    user_id uuid references users(id) on delete cascade,
    book_id uuid references books(id) on delete cascade,
    body text not null,
    primary key (user_id, book_id)
);
