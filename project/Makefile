DICANE_BINARY=dicaneApp

## up: starts all containers in the background without forcing build
up: 
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_dicane
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## build_dicane: builds the dicane binary as a linux executable
build_dicane:
	@echo Building dicane binary...
	chdir ..\dicane-api && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${DICANE_BINARY} ./cmd/api
	@echo Done!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!
