ALTER TABLE cart DROP CONSTRAINT IF EXISTS cart_batch_id_fkey;

ALTER TABLE cart RENAME COLUMN batch_id TO product_id;

ALTER TABLE cart
ADD CONSTRAINT cart_product_id_fkey FOREIGN KEY (product_id) REFERENCES product(product_id);
