db:
	docker compose -f ./infra/docker-compose.yaml up -d
db-stop:
	docker compose -f ./infra/docker-compose.yaml down

db-restart:
	docker compose -f ./infra/docker-compose.yaml down 
	docker compose -f ./infra/docker-compose.yaml up -d

db-clear:
	docker compose -f ./infra/docker-compose.yaml rm -V

