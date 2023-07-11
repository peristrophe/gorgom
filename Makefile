include database/Makefile

export GORGOM_DB_HOST       := 172.56.56.100
export GORGOM_DB_PORT       := 5432
export GORGOM_DB_USER       := test
export GORGOM_DB_PASSWORD   := test
export GORGOM_DB_NAME       := gorgom

.PHONY: run migrate clear-build-cache

run:
	go run cmd/gorgom/gorgom.go

migrate:
	go run cmd/migrator/migrator.go

# out of dev-container
clear-build-cache:
	docker builder prune --force
