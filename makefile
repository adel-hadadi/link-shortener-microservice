LINK_SHORTENER_BINARY=linkShortenerApp
REPORT_BINARY=reportApp
GATEWAY_BINARY=gatewayApp

build-shortener:
	cd ./internal/link && env GOOS=linux CGO_ENABLED=0 go build -o ${LINK_SHORTENER_BINARY} .

build-report:
	cd ./internal/report && env GOOS=linux CGO_ENABLED=0 go build -o ${REPORT_BINARY} .

build-gateway:
	cd ./internal/gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} .

up:
	docker compose up -d

down:
	docker compose down

up-build: build-gateway build-shortener build-report
	@echo "Stop docker containers"
	docker compose down
	@echo "Start docker containers"
	docker compose up -d --build
	@echo "Done"

logs:
	@docker compose logs -f

ps:
	@docker compose ps

openapi-gen:
	@./scripts/openapi-http.sh link internal/link/ports ports
	@./scripts/openapi-http.sh gateway internal/gateway main

proto:
	@./scripts/proto.sh report
	@./scripts/proto.sh link
