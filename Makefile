include database/Makefile

.PHONY: run migrate clear-build-cache

run:
	go run cmd/gorgom/gorgom.go

build:
	go build -o app cmd/gorgom/gorgom.go

migrate:
	go run cmd/migrator/migrator.go

# out of dev-container
clear-build-cache:
	docker builder prune --force
