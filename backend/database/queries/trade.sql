-- name: GetTrade :one
select sqlc.embed(trades), sqlc.embed(users), sqlc.embed(books), sqlc.embed(trade_books) from trades
join users on owner_id = users.id
join users on buyer_id = users.id
join trade_books on id = trade_books.trade_id
join books on book_id = books.id
where trades.id = @id
limit 1;

-- name: ListTrades :many
select sqlc.embed(trades), sqlc.embed(users), sqlc.embed(books), sqlc.embed(trade_books) from trades
join users on owner_id = users.id
join users on buyer_id = users.id
join trade_books on id = trade.trade_id
join books on book_id = book.id
limit $1
offset $2;

-- name: CreateTrade :one
insert into trades (
    owner_id,
    buyer_id,
    created_at
) values (
    @ownerID,
    @buyerID,
    now()
)
returning *;

-- name: AddBookToTrade :exec
insert into trade_books (
    trade_id,
    book_id,
    owner_id
) values (
    @tradeID,
    @bookID,
    @ownerID
);

-- name: CompleteTrade :one
update trades
set completed_at = now()
where id = @id
returning *;

-- name: DeleteTrade :exec
delete from trades
where id = @id;
