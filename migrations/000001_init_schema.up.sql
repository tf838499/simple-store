CREATE TABLE "orders" (
  "id" integer GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY ,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz,
  "delete_at" timestamptz,
  "amount" int[],
  "owner" text,
  "customer_message" text,
  "good_id" text[],
  "total_price" int,
  "message" text,
  "status" int
);

CREATE TABLE "carLists" (
  "id" integer PRIMARY KEY,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz,
  "delete_at" timestamptz,
  "email" text,
  "good_id" int[],
  "good_price" int[],
  "Allimage_name" text
);

CREATE TABLE "goods" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz,
  "delete_at" timestamptz,
  "image_name" text,
  "descript" text,
  "price" bigint,
  "class" text
);

-- ALTER TABLE "orders" ADD FOREIGN KEY ("good_id") REFERENCES "goods" ("id");

-- ALTER TABLE "orders" ADD FOREIGN KEY ("id") REFERENCES "carLists" ("order_id");