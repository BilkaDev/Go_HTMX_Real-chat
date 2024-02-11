run:
	@templ generate
	@cd ./bin/builder & go run staticFile.go
	@go run cmd/main.go cmd/api.go
build:
	@cd ./bin/builder & go run staticFile.go
	@npx tailwindcss -i ./public/assets/css/index.css -o ./public/assets/css/output.css
	@go build -o ./bin/ ./cmd
	@echo "Build complited"

watch:
	@npx tailwindcss -i ./public/assets/css/index.css -o ./public/assets/css/output.css --watch

build-static-file:
	@cd ./bin/builder & go run staticFile.go

build-style:
	@npx tailwindcss -i ./public/assets/css/index.css -o ./public/assets/css/output.css 