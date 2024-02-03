run: build
	@./tmp/main

build: templ tailwind
	@go build -o ./tmp/main .

templ:
	@templ generate

tailwind:
	@npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --minify

watch:
	@air & \
	npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css --watch

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...