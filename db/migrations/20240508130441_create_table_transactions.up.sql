CREATE TABLE transactions(
    transaction_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id),
    product_id INT REFERENCES "product"(product_id),
    batch_id INT REFERENCES "batches"(batch_id),
    transaction_date DATE NOT NULL,
    total_payment NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL,
    transaction_status VARCHAR(12) NOT NULL,
    cancelled_transaction_date DATE NOT NULL,
    cancelled_transaction_reason VARCHAR(255) NOT NULL
);