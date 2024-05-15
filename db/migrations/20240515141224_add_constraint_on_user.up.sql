ALTER TABLE "user"
ADD CONSTRAINT unique_email UNIQUE (email),
ADD CONSTRAINT unique_username UNIQUE (username),
ADD CONSTRAINT unique_whatsapp_number UNIQUE (whatsapp_number);