
.PHONY: prepare run

prepare:
	$(MAKE) prepare -f ./test/Makefile

run:
	go run main.go
