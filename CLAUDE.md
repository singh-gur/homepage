# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web application built with the Gowebly CLI framework. It's a personal homepage/portfolio site using:

- **Backend**: Go 1.24 with built-in net/http package (runs on port 7000)
- **Frontend**: htmx with Alpine.js for reactivity, Tailwind CSS with daisyUI components
- **Templates**: Templ templates for server-side rendering
- **Build Tools**: Parcel for frontend assets, Bun as frontend runtime, Air for live-reloading
- **Module**: `github.com/singh-gur/homepage`

## Common Commands

### Development
```bash
# Start development server with live-reloading (recommended)
gowebly run

# Alternative: Use Air directly for live-reloading
air

# Generate Go files from Templ templates (required before running/building)
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
- `server.go`: HTTP server setup, routing, middleware, and embedded static files
- `handlers.go`: HTTP request handlers for routes and API endpoints
- `templates/`: Templ template files that generate `*_templ.go` files
  - `main.templ`: Base layout template with meta tags, CSS/JS includes
  - `pages/index.templ`: Homepage content with sections for hero, about, projects, experience, education, contact
- `assets/`: Source frontend files (scripts.js, styles.css)
- `static/`: Built frontend assets and static files (served via embed.FS)

### Key Patterns
- **Template Generation**: Templ files must be generated before running: `go run github.com/a-h/templ/cmd/templ@latest generate`
- **Static Files**: Embedded using `//go:embed all:static` in `server.go:14`
- **HTMX Integration**: Uses `github.com/angelofallars/htmx-go` for responses and request validation
- **Routing**: Uses Go 1.22+ enhanced routing patterns (`GET /`, `GET /api/hello-world`)
- **Logging**: Structured logging with `log/slog` throughout request handlers
- **API Endpoints**: Check for HTMX requests using `htmx.IsHTMX(r)` before processing
- **Environment**: Server port configurable via `BACKEND_PORT` env var (defaults to 7000)

### Dependencies
- **Core**: `github.com/a-h/templ`, `github.com/angelofallars/htmx-go`, `github.com/gowebly/helpers`
- **Frontend**: tailwindcss, daisyui, alpinejs, htmx.org
- **Build**: parcel, prettier, postcss
- **Dev Tools**: Air for live-reloading with automatic frontend builds

## Development Workflow

1. **Template Generation**: Always run `go run github.com/a-h/templ/cmd/templ@latest generate` before building/running
2. **Live Development**: Use `gowebly run` or `air` - Air automatically runs `bun run build` before Go builds (.air.toml:8)
3. **Frontend Assets**: Built from `assets/` to `static/` directory via Parcel
4. **File Watching**: Air excludes generated files (`_templ.go`) and static assets from watching
5. **Production**: Ensure Templ files are generated before Docker build

## Request Flow

- `GET /` → `indexViewHandler` → renders homepage with Templ layout
- `GET /api/hello-world` → `showContentAPIHandler` → HTMX-only API endpoint
- `GET /static/*` → served from embedded static files