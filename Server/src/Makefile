## Set ENV by running make ENV=local or develop, staging, or production
ENV := local
CONFIG_FILE ?= ./config/$(ENV).yml
DB_URL ?= $(shell sed -n 's/^dburl:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := migrate -path=db/migrations -database "$(DB_URL)"
LDFLAGS := -ldflags "-X github.com/neelchoudhary/boncuisine/driver.ENV=$(ENV)"

.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-db
run-db: ## run the API server with a dockerized psql server
	docker run --rm   --name boncuisine-db -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres

.PHONY: stop-db
stop-db: ## stop the database
	docker stop boncuisine-db

.PHONY: run-server 
run-server: ## run a local containerized instance of the api
	./entrypoint.sh

.PHONY: run-docker-server
build-docker-server: ## use for testing if builds work
	docker build --build-arg APP_ENV=local -t boncuisine-server .

.PHONY: fake-data
fake-data: ## populate the database with test data
	make migrate-reset
	@echo "Populating test data..."
	@docker exec -it postgres psql "$(DB_URL)" -f /testdata/testdata.sql

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

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