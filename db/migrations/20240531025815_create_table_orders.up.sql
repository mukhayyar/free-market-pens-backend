CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES "transactions"(transaction_id),
    batch_id INT REFERENCES "batches"(batch_id),
    total_payment NUMERIC NOT NULL,
    stock INT NOT NULL
);
