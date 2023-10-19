postgres:
	docker run --name nc-bank-db -p 5432:5432 -e POSTGRES_PASSWORD=pavan -e POSTGRES_USER=pavan -d postgres:12-alpine

createdb:
	docker exec -it nc-bank-db createdb --username=pavan --owner=pavan ncbank

dropdb:
	docker exec -it nc-bank-db dropdb ncbank

migrateup:
	migrate -path db/migration -database "postgresql://pavan:pavan@localhost:5432/ncbank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://pavan:pavan@localhost:5432/ncbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc make