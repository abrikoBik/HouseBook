alter table users add rating real not null default 5 check (rating >= 1 and rating <= 5);

---- create above / drop below ----

alter table users drop column rating;
