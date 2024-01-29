build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./...
	
# adding the at sign ensure that the output is not printed out
