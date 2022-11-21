postgres:
	docker run --name postgres-15 -e POSTGRES_PASSWORD=sY54tji6XqFl -p 5433:5432 -e POSTGRES_USER=root -d postgres:15-alpine
createdb:
	docker exec -it  postgres-15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it  postgres-15 dropdb  simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:sY54tji6XqFl@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:sY54tji6XqFl@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
