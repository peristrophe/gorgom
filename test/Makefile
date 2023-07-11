# It expect run on host machine, not DevContainer

.PHONY: prepare run-postgres

prepare:
ifeq ($(shell docker network ls | grep gorgom_network),)
	docker network create --subnet=172.56.56.0/24 gorgom_network
endif

run-postgres:
	$(MAKE) prepare
	docker run --rm \
	    -e POSTGRES_USER=test \
		-e POSTGRES_PASSWORD=test \
		--net gorgom_network \
		--ip 172.56.56.100 \
		-p 5432:5432 \
		postgres:15.3-bullseye
