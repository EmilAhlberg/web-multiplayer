
clean:
	@go mod tidy

build: clean
	@go build cmd

start: build
	@go run cmd/main.go