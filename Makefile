migrate:
	migrate -path schema/migrations -database postgres://postgres:postgres@localhost:5432/phone-book?sslmode=disable up

create-migration:
	migrate create -ext sql -dir schema/migrations -seq ${name}