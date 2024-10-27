LINK_SHORTENER_BINARY=linkShortenerApp

build-shortener:
	cd ./internal/link && env GOOS=linux CGO_ENABLED=0 go build -o ${LINK_SHORTENER_BINARY} .

up:
	docker compose up -d

up-build: build-shortener
	@echo "Stop docker containers"
	docker compose down
	@echo "Start docker containers"
	docker compose up -d --build
	@echo "Done"

logs:
	@docker compose logs -f

openapi-gen:
	@./scripts/openapi-http.sh link internal/link/ports ports