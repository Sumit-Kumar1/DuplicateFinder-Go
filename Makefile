# run : generate templ and run the project
generate:
	templ generate

run: generate
	@go run cmd/main.go

lint: generate
	go fmt ./...
	golangci-lint run ./...	
