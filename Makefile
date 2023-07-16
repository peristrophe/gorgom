include database/Makefile

TEST_FLAGS := -v

.PHONY: run migrate mock test clear-build-cache

run:
	go run cmd/gorgom/gorgom.go

build:
	go build -o app cmd/gorgom/gorgom.go

migrate:
	go run cmd/migrator/migrator.go

mock:
	-rm -rf ./internal/mock
	go generate ./...

test:
	$(MAKE) mock
	go test $(TEST_FLAGS) ./...

# out of dev-container
clear-build-cache:
	docker builder prune --force
