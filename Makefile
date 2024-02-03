run: templ tailwind
	@go run .

build: templ tailwind
	@go build -o ./tmp/ .

tailwind:
	@npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --minify

templ:
	@templ generate

watch:
	@air & \
	npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --watch

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...