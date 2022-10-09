WEBSH_CLI_BINARY=websh
WEBSH_FRONT_BINARY=webshFrontApp
WEBSH_WS_BINARY=webshWSApp

build-websh-cli:
	@echo "Building ${WEBSH_CLI_BINARY} binary..."
	cd webshell-cli && env GOOS=linux CGO_ENABLED=0 go build -o ${WEBSH_CLI_BINARY} ./cmd/shell && mv ${WEBSH_CLI_BINARY} ../webshell-ws
	@echo "Done!"

build-websh-front:
	@echo "Building ${WEBSH_FRONT_BINARY} binary..."
	cd webshell-front && env GOOS=linux CGO_ENABLED=0 go build -o ${WEBSH_FRONT_BINARY} ./cmd/web
	@echo "Done!"

build-websh-ws: build-websh-cli
	@echo "Building ${WEBSH_WS_BINARY} binary..."
	cd webshell-ws && env GOOS=linux CGO_ENABLED=0 go build -o ${WEBSH_WS_BINARY} ./cmd/api
	@echo "Done!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose build --no-cache
	docker-compose up -d
	@echo "Docker images started!"