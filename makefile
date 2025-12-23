MIGRATIONS_DIR := ./migrations
.PHONY: migrate-up migrate-down migrate-status migrate-create

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-status:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)