CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    store_id INT REFERENCES "store"(store_id),
    name VARCHAR(255) NOT NULL,
    photo BYTEA NOT NULL,
    ready_status VARCHAR(12) NOT NULL,
    stock INT NOT NULL,
    price NUMERIC(10,2) NOT NULL, -- Format uang dalam rupiah, contoh: 10000.00
    pickup_date DATE NOT NULL, -- Tanggal pickup (tanggal/bulan/tahun)
    pickup_time TIME NOT NULL, -- Waktu pickup (jam:menit:detik)
    category_id INT REFERENCES "category"(category_id),
    pickup_place_id INT REFERENCES "store_pickup_place"(store_pickup_place_id)
);
