install:
    @go mod download
    @go mod tidy
    @bun install

build:
    @templ generate
    @go build -ldflags="-s -w" -o .bin/homepage

run:
    @go run github.com/gowebly/gowebly/v3@latest run

