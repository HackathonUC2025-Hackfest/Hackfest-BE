include .env

.PHONY: compose-up compose-down migrate-up migrate-down

migrate-up:
	@docker compose run --rm migrate -path /db/migrations -database "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL)" -verbose up

migrate-down:
	@echo "y" | docker compose run --rm -T migrate -path /db/migrations -database "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL)" -verbose down

compose-up:
	@docker compose up --detach --build

compose-down:
	@docker compose down