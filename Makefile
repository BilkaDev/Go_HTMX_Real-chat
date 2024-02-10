run:
	@templ generate
	@go run cmd/main.go cmd/api.go 
build:
	@go build -o ./bin/ ./cmd