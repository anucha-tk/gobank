include .env

run-dev:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down
exec:
	docker exec -it go_bank_postgres bin/bash
migrate-up:
	migrate -path db/migration -database "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@localhost:5432/$(DATABASE_NAME)?sslmode=disable" --verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@localhost:5432/$(DATABASE_NAME)?sslmode=disable" --verbose down
