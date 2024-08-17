# run : generate templ and run the project
run:
	go run cmd/main.go

lint:
	go fmt ./...
	golangci-lint run ./...	
