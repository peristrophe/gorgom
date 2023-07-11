
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
	echo $(CURDIR)
