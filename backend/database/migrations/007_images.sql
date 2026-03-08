create extension if not exists "uuid-ossp";
alter table books add image_uuid uuid;
alter table users add image_uuid uuid;

---- create above / drop below ----

alter table books drop column image_uuid;
alter table users drop column image_uuid;
