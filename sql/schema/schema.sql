create table users (
    id uuid primary key,
    username varchar(255) not null,
    hashed_password varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);