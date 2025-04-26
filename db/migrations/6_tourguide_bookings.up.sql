CREATE TABLE tourguide_bookings(
    id UUID PRIMARY KEY,
    payment_url VARCHAR NOT NULL,
    star int,
    content TEXT,
    booked_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status VARCHAR NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    tourist_attraction_id UUID REFERENCES tourist_attractions (id) ON DELETE CASCADE
);