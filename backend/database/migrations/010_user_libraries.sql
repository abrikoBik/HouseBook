alter table users add now_reading_id uuid references libraries(id);
alter table users add finished_reading_id uuid references libraries(id);
alter table users add trade_ready_id uuid references libraries(id);

---- create above / drop below ----

alter table users drop column now_reading_id;
alter table users drop column finished_reading_id;
alter table users drop column trade_ready_id;
