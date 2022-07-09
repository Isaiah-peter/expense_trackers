postgres:
	docker run --name postgres13 -p 5342:5432 -e POSTGRES_USER=postgres12 -e POSTGRES_PASSWORD=secret -d postgres:14.3-alpine

createdb:
	docker exec -it postgres13 createdb --username=postgres12 --owner=postgres12 expense_tracker

dropdb:
	docker exec -it postgres13 dropdb expense_tracker -U postgres12

migrateup:
	migrate -path db/migration -database "postgresql://postgres12:secret@localhost:5342/expense_tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres12:secret@localhost:5342/expense_tracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb migratedown migrateup sqlc