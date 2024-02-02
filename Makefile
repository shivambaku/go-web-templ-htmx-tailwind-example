run:
	@templ generate
	@go run cmd/main.go

templ:
	@templ generate