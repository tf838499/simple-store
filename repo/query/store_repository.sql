-- name: GetGoodList :many
SELECT * FROM goods;

-- name: InsertGoods :exec 
INSERT INTO goods (id,image_name,descript,price,class) 
VALUES ($1, $2, $3,$4,$5);

-- name: GetGoodListByPage :many
SELECT * FROM goods ORDER BY id LIMIT $1 OFFSET $2 ;

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