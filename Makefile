postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secrets -p 5432:5432 -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root easy_bank

dropdb:
	docker exec -it postgres dropdb easy_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secrets@localhost:5432/easy_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secrets@localhost:5432/easy_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: postgres createdb dropdb migrateup migratdown sqlc