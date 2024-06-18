db:
	@docker compose up -d
db-stop:
	@docker compose down

db-restart:
	@docker compose down 
	
	@docker compose up -d

#Migrations

cmt: #CREATE TABLE MIGRATION
	@migrate create -ext sql -dir ./cmd/migrations/tables -seq $(filter-out $@,$(MAKECMDGOALS))
cms: #CREATE SEED MIGRATION
	@migrate create -ext sql -dir ./cmd/migrations/seeders -seq $(filter-out $@,$(MAKECMDGOALS))
mu-tables: #UP TABLES
	@go run ./cmd/migrations/main.go up tables
mu-seeders: #UP SEEDERS
	@go run ./cmd/migrations/main.go up seeders
md-tables: #DOWN TABLES
	@go run ./cmd/migrations/main.go down 
# dbforce-version:
# migrate -database "postgres://root:root@localhost:5432/postgres?sslmode=disable" -path cmd/migrations/tables down

#Tests

it: #Integration Tests
	@go test ./test/integration/
