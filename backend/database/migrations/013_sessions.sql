create table sessions (
    id uuid primary key default uuidv7(),
    jwt text not null,
    user_id uuid references users(id) on delete cascade,
);

---- create above / drop below ----

drop table sessions;
