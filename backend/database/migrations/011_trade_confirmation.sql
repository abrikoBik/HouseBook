alter table trades add owner_confirmed boolean not null default false;
alter table trades add buyer_confirmed boolean not null default false;

---- create above / drop below ----

alter table trades drop column owner_confirmed;
alter table trades drop column buyer_confirmed;
