CREATE TABLE store (
    store_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id),
    name VARCHAR(15) NOT NULL,
    photo_profile TEXT NOT NULL,
    whatsapp_number VARCHAR(12) NOT NULL
);