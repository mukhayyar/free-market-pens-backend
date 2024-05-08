CREATE TABLE cart(
    cart_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id),
    product_id INT REFERENCES "product"(product_id),
    quantity INT NOT NULL
);