CREATE TABLE "orders" (
  "id" integer PRIMARY KEY,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz,
  "delete_at" timestamptz,
  "good_id" int,
  "amount" int
);

CREATE TABLE "carLists" (
  "id" integer PRIMARY KEY,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz,
  "delete_at" timestamptz,
  "account" text,
  "phone" text,
  "email" text,
  "total" bigint,
  "status" int,
  "order_id" int
);

CREATE TABLE "goods" (
  "id" integer PRIMARY KEY,
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
