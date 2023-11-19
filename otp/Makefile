my-postgres:
	docker run --name my-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it my-postgres createdb --username=root --owner=root otp

dropdb:
	docker exec -it my-postgres dropdb otp

migrateup:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/otp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/otp?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
server:
	go run main.go


.PHONY: createdb dropdb migrateup migratedown sqlc  server my-postgres