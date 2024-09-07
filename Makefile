ifneq (,$(wildcard ./.env))
    include .env
    export
endif

BINARY_NAME=go-event-api


## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -buildvcs=false -o ${BINARY_NAME} .
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	./${BINARY_NAME}
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...

migrate-up-all:
	@echo "Migrating up...${DATABASE}"
	@migrate -path ./migrations -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_URL})/${DATABASE}" up

migrate-down-all:
	@echo "Migrating down...${DATABASE}"
	@migrate -path ./migrations -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_URL})/${DATABASE}" down

migrate-down-one:
	@echo "Migrating down...${DATABASE}"
	@migrate -path ./migrations -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_URL})/${DATABASE}" down 1

start-prod: migrate-up-all start