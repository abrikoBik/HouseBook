alter table library_books drop column quantities;
alter table library_books add exchangeable boolean not null default false;
alter table trade_books add library_id uuid references libraries(id);

---- create above / drop below ----

alter table library_books add quantity numeric not null default 1;
alter table library_books drop column exchangeable;
alter table trade_books drop column library_id;
