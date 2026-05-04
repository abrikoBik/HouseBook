create extension if not exists "uuid-ossp";
create table libraries (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id),
    name text not null default 'Библиотекa'
);

create table library_books (
    library_id uuid references libraries(id) on delete cascade,
    book_id uuid references books(id),
    quantity numeric not null default 1,
    primary key (library_id, book_id)
);

---- create above / drop below ----

drop table libraries;
drop table library_books;
