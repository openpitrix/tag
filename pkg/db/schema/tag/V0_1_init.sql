
drop table if exists tags;

create table tag
(
   tag_id               varchar(50),
   key         varchar(50),
   value        varchar(50),
   creator       varchar(50),
   create_time          timestamp not null default CURRENT_TIMESTAMP,
   update_time          timestamp not null default CURRENT_TIMESTAMP
);

create table resource_tag
(
   resource_tag_id               varchar(50),
   resource_id         varchar(50),
   tag_id        varchar(50),
   create_time          timestamp not null default CURRENT_TIMESTAMP
);