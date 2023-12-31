DB_URL=postgresql://root:secret@localhost:5432/db_tasks?sslmode=disable

postgres:
	docker run --name pg-local -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg-local createdb --username=root --owner=root db_tasks

dropdb:
	docker exec -it pg-local dropdb db_tasks

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

re-db: dropdb createdb migrateup

sqlc-win:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

compose-up:
	docker-compose up --build -d

compose-down:
	docker-compose down

run:
	go run ./cmd/

.PHONY: postgres createdb migrateup migrateup1 migratedown migratedown1 new_migration re-db run compose-up compose-down
