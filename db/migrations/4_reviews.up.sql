CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    star int NOT NULL,
    content VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    local_id UUID REFERENCES locals (id) ON DELETE CASCADE
);