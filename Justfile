APP := file_name(`grep "^module " go.mod | awk '{print $2}'`)
BIN := "./.bin"
GOBIN := "cmd"
WHAT := "./..."
export CGO_ENABLED := "0"

VERSION := `git describe --tags --exact-match 2>/dev/null || git rev-parse --abbrev-ref HEAD`

DOCKERFILE := "Dockerfile"
DOCKER_IMAGE := "mrehbr/ton-liteserver-prometheus-exporter"
DOCKER_BUILD_OPTS := "--build-arg APP_NAME="+APP

# Show available targets
help:
    @just --list

# Build application binary
build name=APP *opts="-trimpath -ldflags '-s -w'":
    @echo "Building {{name}} to {{clean(join(BIN, name, name))}}"
    @go build {{opts}} -o {{clean(join(BIN, name, name))}} {{clean(join(GOBIN, name, "main.go"))}}

# Install aplication into $GOBIN
[no-cd]
install name=APP *opts="-trimpath -ldflags '-s -w'":
    @echo "Installing {{name}}"
    cd {{clean(join(GOBIN, name))}} && go install {{opts}}

# Generate
generate:
    @go generate -x {{WHAT}}

# Run tests
test *opts="-v -test.timeout=1m -cover":
    @go test {{WHAT}} {{opts}}

# Lint code
lint *opts="-v":
    @golangci-lint run {{opts}} {{WHAT}}

# Format code
fmt:
    @gofumpt -extra -l -w `go list -f '{{{{.Dir}}' {{WHAT}} | rg -v mocks`

# Tidy dependencies
tidy:
    @go mod tidy

# Download dependencies
deps:
    @go mod download


