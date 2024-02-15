run: build
	@./tmp/main

watch:
	@air & \
	npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --watch

build: templ tailwind
	@go build -o ./tmp/main .

templ:
	@templ generate

tailwind:
	@npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --minify

.PHONY: sql
sql:
	@sqlc generate

devdb:
	@docker-compose up -d postgres-dev

devdbdown:
	@docker-compose down

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...