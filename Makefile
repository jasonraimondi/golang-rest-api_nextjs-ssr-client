
DB_CONN := postgres://print:print@localhost:5432/print?sslmode=disable
MIGRATION_DIR := ./db/migrations

MIGRATION := new-migration-name

local:
	docker-compose -f docker-compose.app.yml -f docker-compose.yml up -d

test:
	drone exec --branch master --event push

create-migration:
	 migrate create --dir ${MIGRATION_DIR} --ext sql --seq ${MIGRATION}

migration:
	migrate -source file://${MIGRATION_DIR} -database ${DB_CONN} up

migration-down:
	migrate -source file://${MIGRATION_DIR} -database ${DB_CONN} down 1

migration-down-all:
	migrate -source file://${MIGRATION_DIR} -database ${DB_CONN} down

build:
	docker build -t jasonraimondi/kim-ssr ./web/
	docker push jasonraimondi/kim-ssr
	docker build -t jasonraimondi/kim-api ./
	docker push jasonraimondi/kim-api

.PHONY: dev stage
