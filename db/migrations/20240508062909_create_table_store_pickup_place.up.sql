CREATE TABLE store_pickup_place(
    store_pickup_place_id SERIAL PRIMARY KEY,
    store_id INT REFERENCES "store"(store_id),
    name VARCHAR(255) NOT NULL 
);