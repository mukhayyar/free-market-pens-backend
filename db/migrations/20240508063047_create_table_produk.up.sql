CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    store_id INT REFERENCES "store"(store_id),
    name VARCHAR(255) NOT NULL,
    photo TEXT NOT NULL,
    description TEXT NOT NULL,
    stock INT NOT NULL,
    price NUMERIC(10,2) NOT NULL, -- Format uang dalam rupiah, contoh: 10000.00
    category_id INT REFERENCES "category"(category_id)
);
