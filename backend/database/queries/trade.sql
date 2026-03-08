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

-- name: ConfirmTrade :one
update trades
set
owner_confirmed = case
    when owner_id = @userID then true
    else owner_confirmed
end,
buyer_confirmed = case
    when buyer_id = @userID then true
    else buyer_confirmed
end,
completed_at = case
    when
        (
            (owner_confirmed = true or owner_id = @userID) and
            (buyer_confirmed = true or buyer_id = @userID)
        ) and
        completed_at is null
    then now()
    else completed_at
end
where id = @id
returning *;

-- name: DeleteTrade :exec
delete from trades
where id = @id;

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

-- name: DeleteBookFromTrade :exec
delete from trade_books
where
trade_id = @trade_id and
book_id = @book_id;
