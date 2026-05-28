alter table trades
    drop constraint "trades_owner_id_fkey",
    drop constraint "trades_buyer_id_fkey",
    add constraint "trades_owner_id_fkey" foreign key (owner_id) references users(id) on delete cascade,
    add constraint "trades_owner_buyer_id_fkey" foreign key (buyer_id) references users(id) on delete cascade;
alter table libraries
    drop constraint "libraries_user_id_fkey",
    add constraint "libraries_user_id_fkey" foreign key (user_id) references users(id) on delete cascade;
alter table trade_books
    drop constraint "trade_books_book_id_fkey",
    drop constraint "trade_books_library_id_fkey",
    drop constraint "trade_books_owner_id_fkey",
    add constraint "trade_books_book_id_fkey" foreign key (book_id) references books(id) on delete cascade,
    add constraint "trade_books_library_id_fkey" foreign key (library_id) references libraries(id) on delete cascade,
    add constraint "trade_books_owner_id_fkey" foreign key (owner_id) references users(id) on delete cascade;
alter table library_books
    drop constraint "library_books_book_id_fkey",
    add constraint "library_books_book_id_fkey" foreign key (book_id) references books(id) on delete cascade;

---- create above / drop below ----
alter table trades
    drop constraint "trades_owner_id_fkey",
    drop constraint "trades_buyer_id_fkey",
    add constraint "trades_owner_id_fkey" foreign key (owner_id) references users(id),
    add constraint "trades_buyer_id_fkey" foreign key (buyer_id) references users(id);
alter table libraries
    drop constraint "libraries_user_id_fkey",
    add constraint "libraries_user_id_fkey" foreign key (user_id) references users(id);
alter table trade_books
    drop constraint "trade_books_book_id_fkey",
    drop constraint "trade_books_library_id_fkey",
    drop constraint "trade_books_owner_id_fkey",
    add constraint "trade_books_book_id_fkey" foreign key (book_id) references books(id),
    add constraint "trade_books_library_id_fkey" foreign key (library_id) references libraries(id),
    add constraint "trade_books_owner_id_fkey" foreign key (owner_id) references users(id);
alter table library_books
    drop constraint "library_books_book_id_fkey",
    add constraint "library_books_book_id_fkey" foreign key (book_id) references books(id);
