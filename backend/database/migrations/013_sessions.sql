create table sessions (
    id uuid primary key default uuidv7(),
    name text not null,
    jwt text unique not null,
    user_id uuid references users(id) on delete cascade
);

---- create above / drop below ----

drop table sessions;
