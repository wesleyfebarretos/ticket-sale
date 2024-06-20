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
mup: #MIGRATIONS UP
	@go run ./cmd/migrations/main.go up
mdown: #MIGRATIONS DOWN 
	@go run ./cmd/migrations/main.go down 
# dbforce-version:
# migrate -database "postgres://root:root@localhost:5432/postgres?sslmode=disable" -path cmd/migrations/tables down

#Tests

it: #Integration Tests
	@go test -v ./test/integration/
