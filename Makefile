build:
	@go build -o bin/exec

run: build
	@./bin/exec

