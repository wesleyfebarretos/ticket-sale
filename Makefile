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
	@docker compose down postgres --volumes
	@docker compose up postgres -d

start-kafka:
	@docker compose up zookeeper kafka kafka-ui -d 

restart-kafka:
	@docker compose down zookeeper kafka kafka-ui --volumes
	@docker compose up zookeeper kafka kafka-ui -d 

# Migrations
create-table:
	@migrate create -ext=sql -dir=./internal/api/migrations/tables -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_table)

create-seed:
	@migrate create -ext sql -dir ./internal/api/migrations/seeders -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_seeder)

create-view:
	@migrate create -ext sql -dir ./internal/api/migrations/views -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_view)

create-schema:
	@migrate create -ext sql -dir ./internal/api/migrations/schemas -seq $(shell echo $(filter-out $@,$(MAKECMDGOALS))_schema)

migrations-up:
	@go run ./cmd/migrations/main.go up
	
migrations-down: 
	@go run ./cmd/migrations/main.go down 

# Tests
# Run integration tests
it: 
	@clear
	@go test -v ./internal/api/tests/integration/

generate-swagger-docs:
	@rm -rf ./internal/api/docs
	@swag init -g ./cmd/api/main.go -o ./internal/api/docs/

format-swagger-configs:
	@swag fmt

