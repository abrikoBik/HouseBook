create extension if not exists "uuid-ossp";
create table trades (
    uuid uuid primary key default uuid_generate_v4(),

    created_at timestamp not null default current_timestamp,
    completed_at timestamp
);

---- create above / drop below ----

drop table books;
