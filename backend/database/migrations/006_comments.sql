create table comments (
    id uuid default uuid_generate_v4(),
    user_id bigserial references users(id) on delete cascade,
    book_id uuid references books(id) on delete cascade,
    body text not null,
    primary key (user_id, book_id)
);

---- create above / drop below ----

drop table comments;
