migrateup:
	# migrate -path migrations -database "postgresql://ubuntu:ubuntu@192.168.1.47:5433/aidms_mt?sslmode=disable" -verbose up
	migrate -path migrations -database "postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable" -verbose up
migratedown:
	# migrate -path migrations -database "postgresql://ubuntu:ubuntu@192.168.1.47:5433/aidms_mt?sslmode=disable" -verbose down
	migrate -path migrations -database "postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable" -verbose down
sqlc:
	sqlc generate
makeStore:
	docker build -t e-store .
dbinit:
	migrate create -ext sql -dir ./migrations -seq init_schema 
-- INSERT INTO goods (id,image_name,descript,price,class) 
-- VALUES ($1, $2, $3,$4,$5);