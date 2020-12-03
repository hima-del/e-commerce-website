 CREATE TABLE "customers"("id" bigserial PRIMARY KEY,
 "username" TEXT NOT NULL,
 "password"TEXT NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
 );

 CREATE TABLE "address"("id"bigserial PRIMARY KEY,
 "customer id"INT NOT NULL,
 "billing address" TEXT NOT NULL,
 "shipping address"TEXT NOT NULL UNIQUE,
 "created_at" timestamptz NOT NULL DEFAULT (now())
 );


 CREATE TABLE "orders"("id"bigserial PRIMARY KEY,
 "customer id"INT NOT NULL,
 "order date" DATE NOT NULL,
 "ship date"DATE NOT NULL,
 "shipping address"TEXT NOT NULL,
 "order status"TEXT NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
 );

 CREATE TABLE "order_details"("id"bigserial PRIMARY KEY,
 "product id"INT NOT NULL,
 "order id"INT NOT NULL,
 "order number"INT NOT NULL,
 "price"FLOAT NOT NULL,
 "discount"FLOAT NOT NULL,
 "total"FLOAT NOT NULL,
 "quantity"INT NOT NULL,
 "color"TEXT NOT NULL,
 "size"TEXT NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
 );


 CREATE TABLE "products"("id" bigserial PRIMARY KEY,
 "name"TEXT NOT NULL,
 "price"FLOAT NOT NULL,
 "picture"TEXT NOT NULL,
 "created_at" timestamptz NOT NULL DEFAULT (now())
 );
 

 CREATE TABLE "blacklist"("token"TEXT NOT NULL);

ALTER TABLE "address" ADD FOREIGN KEY ("customer id") REFERENCES "customers"("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer id") REFERENCES "customers"("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product id") REFERENCES "products"("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order id") REFERENCES "orders"("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("shipping address") REFERENCES "address"("shipping address");

