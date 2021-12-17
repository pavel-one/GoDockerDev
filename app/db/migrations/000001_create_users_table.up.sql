create table users
(
    id         serial primary key unique,
    username   varchar unique,
    Name       varchar,
    created_at timestamp not null,
    updated_at timestamp not null
);



