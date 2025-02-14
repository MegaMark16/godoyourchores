# Output binary name
BINARY_NAME=godoyourchores
VERSION=$(shell git describe --tags --always)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Go compiler flags
LDFLAGS=-ldflags="-X 'main.Version=${VERSION}' -X 'main.BuildTime=${BUILD_TIME}'"

# Default build
all: clean build

# Build application
build:
	go build ${LDFLAGS} -o ./bin/${BINARY_NAME} ./cmd/server

# Run application
run: build
	./bin/${BINARY_NAME}

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f ./bin/${BINARY_NAME}
