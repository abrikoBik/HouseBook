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

-- name: ListUserTrades :many
select * from trades
where
owner_id = @userID or
buyer_id = @userID
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

-- name: DeleteTrade :exec
delete from trades
where id = @id;

-- name: AddBookToTrade :exec
insert into trade_books (
    trade_id,
    book_id,
    library_id,
    owner_id
) values (
    @tradeID,
    @bookID,
    @libraryID,
    @ownerID
);

-- name: DeleteBookFromTrade :exec
delete from trade_books
where
trade_id = @trade_id and
book_id = @book_id;

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

-- name: MoveTradeBooks :exec
update library_books lb
set
    library_id = recipient_lib.id,
    exchangeable = false
from trade_books tb
join trades t ON tb.trade_id = t.id
join libraries recipient_lib on recipient_lib.user_id = (
    case
        when tb.owner_id = t.owner_id then t.buyer_id
        else t.owner_id
    end
)
where tb.trade_id = @id
  and t.completed_at is not null
  and lb.library_id = tb.library_id
  and lb.book_id = tb.book_id;
