run: templ
	@go run .

air:
	@air

templ:
	@templ generate

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...