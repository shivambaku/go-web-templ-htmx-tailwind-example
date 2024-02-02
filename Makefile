run:
	@templ generate
	@go run cmd/main.go

templ:
	@templ generate

lint:
	@golangci-lint run

test:
	@go test -v -cover ./...