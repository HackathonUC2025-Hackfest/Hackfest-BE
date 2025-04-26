CREATE TABLE tourist_attractions (
    id UUID PRIMARY KEY,
    name varchar NOT NULL,
    description text NOT NULL,
    address varchar NOT NULL,
    city varchar NOT NULL,
    province varchar NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    photo_url text NOT NULL,
    tour_guide_price BIGINT NOT NULL,
    tour_guide_count int NOT NULL,
    tour_guide_discount_percentage NUMERIC(5,2) NOT NULL,
    price BIGINT NOT NULL,
    discount_percentage NUMERIC(5,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);