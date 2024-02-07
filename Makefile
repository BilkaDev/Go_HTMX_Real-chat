run:
	@templ generate
	@go run cmd/main.go cmd/api.go cmd/config.go cmd/storage.go