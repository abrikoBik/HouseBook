create table comments (
    id uuid primary key default uuid_generate_v4(),
    user_id bigserial references users(id) on delete cascade,
    book_id uuid references books(id) on delete cascade,
    body text not null
);

---- create above / drop below ----

drop table comments;
