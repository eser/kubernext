.PHONY: build multiarch-build run run-docker clean test cov dep lint
# .RECIPEPREFIX := $(.RECIPEPREFIX)<space>
BINARY_NAME=go-service

build:
	go build -o ./${BINARY_NAME} ./cmd/${BINARY_NAME}

multiarch-build:
	GOARCH=amd64 GOOS=darwin go build -o ./${BINARY_NAME}-darwin-amd64 ./cmd/${BINARY_NAME}
	GOARCH=amd64 GOOS=linux go build -o ./${BINARY_NAME}-linux-amd64 ./cmd/${BINARY_NAME}
	GOARCH=amd64 GOOS=windows go build -o ./${BINARY_NAME}-windows-amd64 ./cmd/${BINARY_NAME}
	GOARCH=arm64 GOOS=darwin go build -o ./${BINARY_NAME}-darwin-arm64 ./cmd/${BINARY_NAME}
	GOARCH=arm64 GOOS=linux go build -o ./${BINARY_NAME}-linux-arm64 ./cmd/${BINARY_NAME}
	GOARCH=arm64 GOOS=windows go build -o ./${BINARY_NAME}-windows-arm64 ./cmd/${BINARY_NAME}

run: build
	./${BINARY_NAME}

run-docker:
	docker compose -f ./docker-compose.yml up

clean:
	go clean

test:
	go test -v ./...

cov:
	go test -v --coverpkg=./... --coverprofile=./tmp/coverage.out ./...
	`go env GOPATH`/bin/gcov2lcov -infile ./tmp/coverage.out -outfile ./tmp/coverage.lcov

dep:
	go mod download
	go mod tidy

lint:
	`go env GOPATH`/bin/golangci-lint run
