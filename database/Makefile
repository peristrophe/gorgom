# It expect run on host machine, not DevContainer

.PHONY: prepare run-postgres

prepare:
ifeq ($(shell docker network ls | grep gorgom_network),)
	docker network create --subnet=172.56.56.0/24 gorgom_network
endif

# for Operation in Dev-Container
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
