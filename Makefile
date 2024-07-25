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
	@migrate create -ext=sql -dir=./api/migrations/tables -seq $(filter-out $@,$(MAKECMDGOALS))

create-seed:
	@migrate create -ext sql -dir ./api/migrations/seeders -seq $(filter-out $@,$(MAKECMDGOALS))

create-view:
	@migrate create -ext sql -dir ./api/migrations/views -seq $(filter-out $@,$(MAKECMDGOALS))

create-schema:
	@migrate create -ext sql -dir ./api/migrations/schemas -seq $(filter-out $@,$(MAKECMDGOALS))

migrations-up:
	@go run ./api/migrations/main.go up
	
migrations-down: 
	@go run ./api/migrations/main.go down 

# Tests
# Run integration tests
it: 
	@clear
	@go test -v ./api/tests/integration/

generate-swagger-docs:
	@rm -rf ./api/docs
	@swag init -g ./cmd/api/main.go -o ./api/docs/

format-swagger-configs:
	@swag fmt

