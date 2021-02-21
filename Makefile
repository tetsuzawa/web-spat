# https://docs.docker.com/docker-for-mac/networking/#use-cases-and-workarounds
DOCKER_DB_DNS:=db

DB_USER:=root
DB_PASSWORD:=password
DB_HOST:=$(DOCKER_DB_DNS)
DB_PORT:=3306
DB_NAME:=web_spat

FLYWAY_CONF?=-url=jdbc:mariadb://$(DOCKER_DB_DNS):$(DB_PORT)/$(DB_NAME) -user=$(DB_USER) -password=$(DB_PASSWORD)

docker-compose/build:
	docker-compose build

docker-compose/up:
	docker-compose up -d

docker-compose/up/service:
	docker-compose up $(service)

docker-compose/down:
	docker-compose down

docker-compose/logs:
	docker-compose logs -f

DB_SERVICE:=db
mariadb/client:
	docker-compose exec $(DB_SERVICE) mysql -u $(DB_USER) -h localhost -p $(DB_PASSWORD) $(DB_NAME)

mariadb/init:
	docker-compose exec $(DB_SERVICE) \
		mysql -u $(DB_USER) -h localhost -p$(DB_PASSWORD) \
		-e "create database \`$(DB_NAME)\`"

__mariadb/drop:
	docker-compose exec $(DB_SERVICE) \
		mysql -u $(DB_USER) -h localhost -p$(DB_PASSWORD) \
		-e "drop database \`$(DB_NAME)\`"

MIGRATION_SERVICE:=migration
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

flyway/validate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) validate

flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline

__flyway/clean:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) clean


BACKEND_SERVICE:=backend
backend/shell:
	docker-compose exec $(BACKEND_SERVICE) bash


openapi/gen/backend:
	@mkdir -p ./backend/interfaces/server/gen
	oapi-codegen -generate "types" docs/openapi.yaml > ./backend/interfaces/server/gen/types.gen.go
	oapi-codegen -generate "server" docs/openapi.yaml > ./backend/interfaces/server/gen/server.gen.go
	oapi-codegen -generate "spec" docs/openapi.yaml > ./backend/interfaces/server/gen/spec.gen.go
