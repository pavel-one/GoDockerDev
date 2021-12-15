create table users
(
    id         serial primary key unique,
    username   varchar unique,
    name       varchar,
    created_at timestamp not null,
    updated_at timestamp not null
);



