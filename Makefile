# Application Services
start-services:
	@chmod +x ./setup.sh
	@./setup.sh

stop-services:
	@docker compose down

kill-services:
	@docker compose down
	@docker volume rm -f $(shell docker volume ls --filter name=tick -q)

rebuild-services: kill-services start-services

restart-db:
	@docker compose down postgres
	@docker volume rm ticket-sale_pgdata
	@docker compose up postgres -d

# Migrations
create-table:
	@migrate create -ext=sql -dir=./cmd/migrations/tables -seq $(filter-out $@,$(MAKECMDGOALS))

create-seed:
	@migrate create -ext sql -dir ./cmd/migrations/seeders -seq $(filter-out $@,$(MAKECMDGOALS))

create-view:
	@migrate create -ext sql -dir ./cmd/migrations/views -seq $(filter-out $@,$(MAKECMDGOALS))

create-schema:
	@migrate create -ext sql -dir ./cmd/migrations/schemas -seq $(filter-out $@,$(MAKECMDGOALS))

migrations-up:
	@go run ./cmd/migrations/main.go up
	
migrations-down: 
	@go run ./cmd/migrations/main.go down 

# Tests
# Run integration tests
it: 
	@clear
	@go test -v ./test/integration/

# Swagger
gen-swg-conf:
	@rm -rf swagger
	@swag init -g ./cmd/main.go -o ./swagger

swg-format:
	@swag fmt

