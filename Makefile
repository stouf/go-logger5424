help:
	@printf "Available commands:\n\n"
	@printf "\tbuild\t\t\tCompile the project\n"
	@printf "\ttest\t\t\tRun the tests\n"
	@printf "\ttest-coverage\t\tRun the tests and display the coverage\n"

build:
	@go build -race ./...

test: build
	@go test -race -cover ./...

test-coverage: build
	@go test -coverprofile=test-coverage.out ./...
	@go tool cover -html=test-coverage.out
	@rm test-coverage.out