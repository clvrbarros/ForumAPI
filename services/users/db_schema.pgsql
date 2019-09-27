CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT,
    country TEXT,
    active BOOLEAN NOT NULL,
    created_at timestamp default current_timestamp
);