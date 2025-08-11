# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web application built with the Gowebly CLI framework. It's a homepage project using:

- **Backend**: Go 1.24 with built-in net/http package (runs on port 7000)
- **Frontend**: htmx with Alpine.js for reactivity, Tailwind CSS with daisyUI components
- **Templates**: Templ templates for server-side rendering
- **Build Tools**: Parcel for frontend assets, Bun as frontend runtime, Air for live-reloading

## Common Commands

### Development
```bash
# Start development server with live-reloading (recommended)
gowebly run

# Alternative: Use Air directly for live-reloading
air

# Generate Go files from Templ templates (required before Docker builds)
go run github.com/a-h/templ/cmd/templ@latest generate
```

### Frontend Build
```bash
# Development build (unoptimized)
bun run dev

# Production build
bun run build

# Watch mode for development
bun run watch

# Format frontend code
bun run fmt
```

### Backend
```bash
# Run the application directly
go run .

# Build binary
go build -o ./tmp/gowebly_default .

# Lint (if golangci-lint is configured)
golangci-lint run
```

### Docker
```bash
# Build and run with Docker Compose
docker-compose up

# Important: Generate Templ files before Docker build
go run github.com/a-h/templ/cmd/templ@latest generate
```

## Architecture

### Project Structure
- `main.go`: Entry point, calls `runServer()`
- `server.go`: HTTP server setup, routing, and middleware
- `handlers.go`: HTTP request handlers for routes and API endpoints
- `templates/`: Templ template files
  - `main.templ`: Base layout template
  - `pages/`: Page-specific templates (index.templ, info.templ)
- `assets/`: Source frontend files (scripts.js, styles.css)
- `static/`: Built frontend assets and static files (served via embed.FS)

### Key Patterns
- Templates use Templ syntax and generate corresponding `*_templ.go` files
- Static files are embedded using `//go:embed all:static` in server.go
- HTMX integration via `github.com/angelofallars/htmx-go` package
- Error handling with structured logging via `log/slog`
- API endpoints check for HTMX requests using `htmx.IsHTMX(r)`

### Dependencies
- Core: `github.com/a-h/templ`, `github.com/angelofallars/htmx-go`, `github.com/gowebly/helpers`
- Frontend: tailwindcss, daisyui, alpinejs, htmx.org
- Build: parcel, prettier, postcss

## Development Workflow

1. Templates must be generated before running: `go run github.com/a-h/templ/cmd/templ@latest generate`
2. Frontend assets are built from `assets/` to `static/` directory
3. Use `gowebly run` or `air` for development with live-reloading
4. Air automatically runs `bun run build` before Go builds (see .air.toml)
5. For production deployment, ensure Templ files are generated before Docker build