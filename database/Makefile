
export GORGOM_DB_HOST       := 172.56.56.100
export GORGOM_DB_PORT       := 5432
export GORGOM_DB_USER       := test
export GORGOM_DB_PASSWORD   := test
export GORGOM_DB_NAME       := gorgom

.PHONY: prepare run-postgres init-postgres

# out of dev-container
prepare:
ifeq ($(shell which docker && docker network ls | grep gorgom_network),)
	docker network create --subnet=172.56.56.0/24 gorgom_network
endif

# out of dev-container
run-postgres:
	$(MAKE) prepare
	docker run --rm \
	    -e POSTGRES_USER=${GORGOM_DB_USER} \
		-e POSTGRES_PASSWORD=${GORGOM_DB_PASSWORD} \
		-e POSTGRES_INITDB_ARGS="--encoding=UTF-8 --locale=C" \
		--net gorgom_network \
		--ip ${GORGOM_DB_HOST} \
		-p ${GORGOM_DB_PORT}:${GORGOM_DB_PORT} \
		postgres:15.3-bullseye

init-postgres:
	-test -f ./init.dcl && psql service=init-dev < ./init.dcl
	-test -f ./database/init.dcl && psql service=init-dev < ./database/init.dcl
	-test -f ./init.dcl && go run ../cmd/migrator/migrator.go
	-test -f ./database/init.dcl && go run cmd/migrator/migrator.go
