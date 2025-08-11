# Project Context for Qwen Code

## Project Overview

This is a web application project generated using the **Gowebly CLI**. It's a full-stack application combining Go for the backend with HTMX, Alpine.js, Tailwind CSS, and DaisyUI for the frontend.

*   **Backend:** Go (using the built-in `net/http` package).
*   **Frontend:** HTMX for interactivity, Alpine.js for complex reactivity, Tailwind CSS + DaisyUI for styling. Templates are generated using the `Templ` tool.
*   **Build Tools:** Bun (runtime), Parcel (bundler), Prettier (formatter).
*   **Development Tools:** Air (live reloading), golangci-lint (linter).

The application serves a simple page with a button that triggers an HTMX request to an API endpoint.

## Key Files and Directories

*   `main.go`: The entry point of the Go application.
*   `server.go`: Configures and starts the HTTP server, handles routing, and serves static files.
*   `handlers.go`: Contains the HTTP handler functions for the index page and API endpoints.
*   `go.mod`, `go.sum`: Go module definition and dependencies.
*   `package.json`: Defines frontend dependencies (Tailwind, DaisyUI, Alpine.js, HTMX) and build scripts.
*   `assets/`: Contains source CSS (`styles.css`) and JavaScript (`scripts.js`) files.
*   `templates/`: Contains `*.templ` files used by the `Templ` tool to generate Go code for HTML rendering.
*   `static/`: Contains built/minified frontend assets (CSS, JS) and static files like images and favicons. This folder is populated by the build process.
*   `.air.toml`: Configuration for the Air live-reloading tool.
*   `Dockerfile`, `docker-compose.yml`: Configuration for containerizing and deploying the application.

## Building and Running

### Prerequisites

Ensure the following tools are installed:

*   `Air`: [https://github.com/air-verse/air](https://github.com/air-verse/air)
*   `Bun`: [https://github.com/oven-sh/bun](https://github.com/oven-sh/bun)
*   `Templ`: [https://github.com/a-h/templ](https://github.com/a-h/templ)
*   `golangci-lint`: [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

### Development

The recommended way to start the development server with live-reloading is using the Gowebly CLI:

```bash
gowebly run
```

This command typically handles running `Air` for the Go backend and `Bun`/`Parcel` for the frontend assets.

Alternatively, you can run the Go server directly (after building templates with `templ generate`):

```bash
go run .
```

To manually build frontend assets (styles and scripts) into the `static/` directory:

```bash
# For development (unoptimized)
bun run dev
# OR for production (optimized)
bun run build
```

To watch for changes in frontend assets and rebuild automatically (useful without `gowebly run` or `air`):

```bash
bun run watch
```

### Formatting

To format frontend files using Prettier:

```bash
bun run fmt
```

### Linting

To lint Go code:

```bash
golangci-lint run
```

## Deployment

Deployment is configured using Docker. The steps are:

1.  Ensure Go code from `*.templ` templates is generated: `templ generate`.
2.  Build frontend assets: `bun run build`.
3.  Use `docker-compose up` to build and run the application based on the `Dockerfile` and `docker-compose.yml`.