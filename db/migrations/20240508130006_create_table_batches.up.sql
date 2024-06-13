CREATE TABLE batches (
    batch_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES "product"(product_id),
    store_pickup_place_id INT REFERENCES "store_pickup_place"(store_pickup_place_id),
    stock INT NOT NULL,
    close_order_time TIMESTAMP,
    pickup_time TIMESTAMP
);
