## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool sokudo and copies it to sokudo-helper
build_cli:
	@go build -o ../myapp/sokudo ./cmd/cli

## build: builds the command line tool dist directory
build:
	@go build -o ./dist/sokudo ./cmd/cli