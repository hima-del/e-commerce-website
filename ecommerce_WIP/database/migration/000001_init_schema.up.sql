 CREATE TABLE "customers"("id" bigserial PRIMARY KEY,
 "username" TEXT NOT NULL,
 "password"TEXT NOT NULL,
 "createdat" timestamptz NOT NULL DEFAULT (now())
 );

 CREATE TABLE "address"("id"bigserial PRIMARY KEY,
 "customerid"INT NOT NULL,
 "billingaddress" TEXT NOT NULL,
 "shippingaddress"TEXT NOT NULL UNIQUE,
 "createdat" timestamptz NOT NULL DEFAULT (now())
 );


 CREATE TABLE "orders"("id"bigserial PRIMARY KEY,
 "customerid"INT NOT NULL,
 "orderdate" DATE NOT NULL,
 "shipdate"DATE NOT NULL,
 "shippingaddress"TEXT NOT NULL,
 "orderstatus"TEXT NOT NULL,
 "createdat" timestamptz NOT NULL DEFAULT (now())
 );

 CREATE TABLE "orderdetails"("id"bigserial PRIMARY KEY,
 "productid"INT NOT NULL,
 "orderid"INT NOT NULL,
 "ordernumber"serial NOT NULL,
 "price"FLOAT NOT NULL,
 "discount"FLOAT NOT NULL,
 "total"FLOAT NOT NULL,
 "quantity"INT NOT NULL,
 "color"TEXT NOT NULL,
 "size"TEXT NOT NULL,
 "createdat" timestamptz NOT NULL DEFAULT (now())
 );


 CREATE TABLE "products"("id" bigserial PRIMARY KEY,
 "name"TEXT NOT NULL,
 "price"FLOAT NOT NULL,
 "picture"TEXT NOT NULL,
 "createdat" timestamptz NOT NULL DEFAULT (now())
 );
 

 CREATE TABLE "blacklist"("token"TEXT NOT NULL);

ALTER TABLE "address" ADD FOREIGN KEY ("customerid") REFERENCES "customers"("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customerid") REFERENCES "customers"("id");

ALTER TABLE "orderdetails" ADD FOREIGN KEY ("productid") REFERENCES "products"("id");

ALTER TABLE "orderdetails" ADD FOREIGN KEY ("orderid") REFERENCES "orders"("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("shippingaddress") REFERENCES "address"("shippingaddress");

