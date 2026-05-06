alter table users
    alter column id drop default,
    alter column id set default uuidv7();
alter table books
    alter column id drop default,
    alter column id set default uuidv7();
alter table trades
    alter column id drop default,
    alter column id set default uuidv7();
alter table libraries
    alter column id drop default,
    alter column id set default uuidv7();
alter table comments
    alter column id drop default,
    alter column id set default uuidv7();

---- create above / drop below ----

create extension if not exists "uuid-ossp";
alter table users
    alter column id drop default,
    alter column id set default uuid_generate_v4();
alter table books
    alter column id drop default,
    alter column id set default uuid_generate_v4();
alter table trades
    alter column id drop default,
    alter column id set default uuid_generate_v4();
alter table libraries
    alter column id drop default,
    alter column id set default uuid_generate_v4();
alter table comments
    alter column id drop default,
    alter column id set default uuid_generate_v4();    
