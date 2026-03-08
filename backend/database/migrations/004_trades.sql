alter table trades add owner_id bigserial not null references users(id);
alter table trades add buyer_id bigserial not null references users(id);

create table trade_books (
    trade_id uuid references trades(id) on delete cascade,
    book_id uuid references books(id),
    owner_id bigserial not null references users(id),
    primary key (trade_id, book_id)
);

---- create above / drop below ----

alter table trades drop column owner_id;
alter table trades drop column buyer_id;

drop table trade_books;
