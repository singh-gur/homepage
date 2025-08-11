package main

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	gowebly "github.com/gowebly/helpers"
)

//go:embed all:static
var static embed.FS

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "7000"))
	if err != nil {
		return err
	}

	// Handle static files from the embed FS (with a custom handler).
	http.Handle("GET /static/", gowebly.StaticFileServerHandler(http.FS(static)))

	// Handle index page view.
	http.HandleFunc("GET /", indexViewHandler)
	
	// Handle about page view.
	http.HandleFunc("GET /about", aboutViewHandler)
	
	// Handle projects page view.
	http.HandleFunc("GET /projects", projectsViewHandler)
	
	// Handle info page view (keep for backward compatibility).
	http.HandleFunc("GET /info", infoViewHandler)

	// Handle API endpoints.
	http.HandleFunc("GET /api/hello-world", showContentAPIHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// Note: The ReadTimeout and WriteTimeout settings may interfere with SSE (Server-Sent Event) or WS (WebSocket) connections.
	// For SSE or WS, these timeouts can cause the connection to reset after 10 or 5 seconds due to the ReadTimeout and WriteTimeout setting.
	// If you plan to use SSE or WS, consider commenting out or removing the ReadTimeout and WriteTimeout key-value pairs.
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
