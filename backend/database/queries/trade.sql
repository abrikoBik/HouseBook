-- name: GetTrade :one
select trades.* from trades
join users as owner on trades.owner_id = owner.id
join users as buyer on trades.buyer_id = buyer.id
left join trade_books on trades.id = trade_books.trade_id
left join books on trade_books.book_id = books.id
where trades.id = @id
limit 1;

-- name: ListTrades :many
select * from trades
limit $1
offset $2;

-- name: ListTradeBooks :many
select book_id, owner_id from trade_books
where trade_id = @id
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
