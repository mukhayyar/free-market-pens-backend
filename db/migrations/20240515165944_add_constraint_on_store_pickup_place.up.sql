ALTER TABLE "store_pickup_place"
ADD CONSTRAINT unique_store_pickup_place UNIQUE (store_id, name);
