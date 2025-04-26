CREATE TABLE users (
    id UUID PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR,
    auth_provider VARCHAR NOT NULL,
    photo_url TEXT NOT NULL,
    is_premium bool not null default false,
    expired_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
