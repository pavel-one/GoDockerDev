create table user_tokens
(
    id          bigserial primary key unique,
    user_id     bigserial not null,
        constraint fk_user_token_user_id
            FOREIGN KEY (user_id)
                REFERENCES users(id)
                    ON DELETE CASCADE,
    token       varchar not null unique,
    created_at      timestamp not null,
    updated_at      timestamp not null,
    latest_use  timestamp not null
)
