postgres:
	docker run --name postgres-15 -e POSTGRES_PASSWORD=devsecret -p 5432:5432 -e POSTGRES_USER=root -d postgres:15-alpine
createdb:
	docker exec -it  postgres-15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it  postgres-15 dropdb  simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:devsecret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:devsecret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
