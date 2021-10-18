.PHONY: docker-start docker-stop start gendoc test
.DEFAULT_GOAL := help
PROJECTNAME := $(shell basename "$(PWD)")

## setup: Initialize project
setup: copy-env docker-start mongo-restore

## docker-start: Start docker-compose
docker-start:
	docker-compose up --build -d

## docker-stop: Stop docker-compose
docker-stop:
	docker-compose down

## start: Start the application
start:
	air

## copy-env: Copy environment file
copy-env:
	cp .env.example .env

## gendoc: Generate docs api with swagger
gendoc:
	swag init -g api/app.go

## test: Run tests coverage
test:
	go test -race -v -cover -covermode=atomic ./...

## mongo-dump: Dump MongoDB data for testing
mongo-dump:
	@echo "Dump MongoDB data for testing"
	@bash seed/mongodb/dump.sh
	@tar -czvf seed/mongodb/dump.tar.gz seed/mongodb/dump
	@rm -r seed/mongodb/dump

## mongo-restore: Restore MongoDB data for testing 
mongo-restore:
	@tar -xzvf seed/mongodb/dump.tar.gz seed/mongodb/dump
	@echo "Import MongoDB data for testing"
	@bash seed/mongodb/restore.sh
	@rm -r seed/mongodb/dump


help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo