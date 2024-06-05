ALTER TABLE cart DROP CONSTRAINT IF EXISTS cart_product_id_fkey;

ALTER TABLE cart RENAME COLUMN product_id TO batch_id;

ALTER TABLE cart
ADD CONSTRAINT cart_batch_id_fkey FOREIGN KEY (batch_id) REFERENCES batches(batch_id);
