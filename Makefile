postgres:
	docker run --name postgres-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres-alpine createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-alpine dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

rootdb:
	docker exec -it postgres-alpine psql -U root

simple_bank_db:
	docker exec -it postgres-alpine psql -U root simple_bank

sqlc:
	docker run --rm -v C:\Users\Lenovo\go\src\go-bank:/src -w /src kjconroy/sqlc:1.4.0 generate

test:
	go test -v -cover ./...

test-with-output:
	go test -coverprofile=coverage.out ./...

show-test-output:
	go tool cover -html=coverage.out

.PHONY: postgres createdb dropdb migrateup migratedown sqlc