-- name: GetGoodList :many
SELECT * FROM goods;

-- name: Insertgoods :exec
INSERT INTO goods (id,image_name,descript,price,class) 
VALUES ($1, $2, $3,$4,$5);
-- CREATE TABLE "goods" (
--   "id" integer PRIMARY KEY,
--   "created_at" timestamptz DEFAULT 'now()',
--   "updated_at" timestamptz,
--   "delete_at" timestamptz,
--   "image_name" text,
--   "desc" text,
--   "price" bigint,
--   "class" text
-- );