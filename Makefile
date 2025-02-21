build:
	@go build -o bin/storages
run : build
	@./bin/storages
test:
	@go test ./.. -v
