alter table books add added_at timestamp not null default current_timestamp;

---- create above / drop below ----

alter table books drop column added_at;
