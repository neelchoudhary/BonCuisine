## Set ENV by running make ENV=local, develop, staging, production
ENV := local
# DB URL is always local
DB_URL := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path=db/migrations -database "$(DB_URL)"

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-server 
run-server: ## run the grpc server locally
	./entrypoint.sh

.PHONY: run-auth-client 
run-auth-client: ## run a local instance of a test client
	go run cmd/mock_client/client.go -env $(ENV) -c auth

.PHONY: run-recipe-client 
run-recipe-client: ## run a local instance of a test client
	go run cmd/mock_client/client.go -env $(ENV) -c recipe

.PHONY: build-docker-server
build-docker-server: ## use for testing if builds work
	docker build --build-arg APP_ENV=$(ENV) -t boncuisine-server .

.PHONY: start-db
start-db: ## start the API server with a dockerized psql server
	docker run --rm   --name boncuisine-db -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres
	docker cp db/sampledata/. boncuisine-db:/var/sampledata

.PHONY: stop-db
stop-db: ## stop the database
	docker stop boncuisine-db

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-force
migrate-force: ## force migration
	@read -p "Forcing database migration at: " name; \
	$(MIGRATE) force $${name// /_}

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir db/migrations -seq $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-test
migrate-test: ## run migrations, drop migrations, then run migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up
	@$(MIGRATE) down
	@$(MIGRATE) up
	@echo "Success..."