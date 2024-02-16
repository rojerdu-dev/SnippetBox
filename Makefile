BINARY_NAME=snippetbox_api

build:
	@GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME} ./cmd/web/

run: build
	@./bin/${BINARY_NAME}

test:
	@go test -v ./...

lint:
	@golangci-lint run

clean:
	@go clean
	@rm bin/${BINARY_NAME}