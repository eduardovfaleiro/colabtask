DB_URL = postgresql://root:password@db:5432/colabtask_db?sslmode=disable
MIGRATE_IMAGE = migrate/migrate
MIGRATION_DIR = db/migration
NETWORK_NAME = $$(basename $$(pwd))_default

postgres:
	docker compose up -d

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root colabtask_db

dropdb:
	docker exec -it postgres18 dropdb colabtask_db

migrateup:
	docker run --rm -v $(shell pwd)/$(MIGRATION_DIR):/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL)" up

# TODO(testar)
migratedown:
	docker run --rm -v $(shell pwd)/$(MIGRATION_DIR):/migrations --network $(NETWORK_NAME) $(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL)" down 1

.PHONY: postgres createdb dropdb migrateup migratedown