-- name: GetGoodList :many
SELECT * FROM goods;

-- name: InsertGoods :exec 
INSERT INTO goods (image_name,descript,price,class) 
VALUES ($1, $2, $3,$4);

-- name: GetGoodListByPage :many
SELECT * FROM goods ORDER BY id LIMIT $1 OFFSET $2 ;

-- name: GetGoodByName :one
SELECT * FROM goods WHERE image_name = $1 LIMIT 1 ;

-- name: UpdateGood :exec
Update goods
SET
    image_name = $1,
    descript = $2,
    price = $3,
    class = $4
WHERE
    id = $5;

-- name: DeleteGood :exec
DELETE FROM goods WHERE id = $1;

-- name: InsertOrder :exec 
INSERT INTO orders (amount,owner,good_id,total_price,message,status) 
VALUES ($1, $2, $3,$4,$5,$6);

-- name: GetGetOrderByOwner :many
SELECT * FROM orders WHERE owner = $1;