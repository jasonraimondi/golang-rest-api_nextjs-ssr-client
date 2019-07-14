
DB_CONN := postgres://print:print@localhost:5432/print?sslmode=disable
MIGRATION_DIR := ./db/migrations

MIGRATION := new-migration-name

test:
	drone exec --branch master --event push

create-migration:
	 migrate create --dir ${MIGRATION_DIR} --ext sql --seq ${MIGRATION}

migration:
	migrate -source file://${MIGRATION_DIR} -database ${DB_CONN} up

migration-down:
	migrate -source file://${MIGRATION_DIR} -database ${DB_CONN} down

.PHONY: dev stage
