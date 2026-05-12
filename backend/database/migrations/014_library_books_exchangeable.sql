alter table library_books add exchangeable boolean not null default false;

---- create above / drop below ----

alter table library_books drop column exchangeable;
