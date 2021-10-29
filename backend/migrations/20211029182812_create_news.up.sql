CREATE TABLE news (
    id bigserial not null primary key,
    title varchar(255) not null,
    created_at timestamp not null
);