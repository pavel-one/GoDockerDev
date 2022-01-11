create table users
(
    id              bigserial primary key unique,
    username        varchar unique,
    name            varchar,
    email           varchar unique,
    password        varchar not null,
    created_at      timestamp not null,
    updated_at      timestamp not null
);
CREATE INDEX idx_username_email ON users(username, email)


