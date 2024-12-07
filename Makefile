build:
		go build -o bob cmd/cli/main.go

start: build
		mkdir -p ../dummy
		mv bob ../dummy/

# Constants
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_USER ?= postgres
DATABASE_PASSWORD ?= postgres
DATABASE_SSL ?= disable
DATABASE_DATABASE = kickstart-db
DATABASE_DSN := "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DATABASE}?sslmode=${DATABASE_SSL}"
MIGRATIONS_DIR="./db/migration"

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $(NAME)