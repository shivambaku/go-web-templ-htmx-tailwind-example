include .env

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

db:
	@docker-compose up -d postgres-dev

db-down:
	@docker-compose down

sqlc:
	@sqlc generate

migrate-diff: sqlc
	@atlas migrate diff --dir "file://sql/migrations" --to "file://sql/schema.sql" --dev-url "docker://postgres/16/dev?search_path=public"

migrate-apply: migrate-diff
	@atlas migrate apply --dir "file://sql/migrations" --url ${DATABASE_URL}

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...
