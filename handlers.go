package main

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"

	"github.com/singh-gur/homepage/templates"
	"github.com/singh-gur/homepage/templates/pages"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"Gurbakhshish Singh, software engineer, portfolio, Go developer, web development",
		"Software Engineer specializing in Go, JavaScript, and modern web technologies. View my portfolio and projects.",
	)

	// Define template body content.
	bodyContent := pages.HomeBodyContent()

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Gurbakhshish Singh - Software Engineer", // define title text
		metaTags,                                 // define meta tags
		bodyContent,                              // define body content
	)

	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, indexTemplate); err != nil {
		// If not, return HTTP 400 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// aboutViewHandler handles a view for the about/resume page.
func aboutViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.AboutMetaTags()

	// Define template body content.
	bodyContent := pages.AboutBodyContent()

	// Define template layout for about page.
	aboutTemplate := templates.Layout(
		"About - Gurbakhshish Singh", // define title text
		metaTags,                     // define meta tags
		bodyContent,                  // define body content
	)

	// Render about page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, aboutTemplate); err != nil {
		// If not, return HTTP 500 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// projectsViewHandler handles a view for the projects page.
func projectsViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.ProjectsMetaTags()

	// Define template body content.
	bodyContent := pages.ProjectsBodyContent()

	// Define template layout for projects page.
	projectsTemplate := templates.Layout(
		"Projects - Gurbakhshish Singh", // define title text
		metaTags,                        // define meta tags
		bodyContent,                     // define body content
	)

	// Render projects page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, projectsTemplate); err != nil {
		// If not, return HTTP 500 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// infoViewHandler handles a view for the info page.
func infoViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.InfoMetaTags()

	// Define template body content.
	bodyContent := pages.InfoBodyContent()

	// Define template layout for info page.
	infoTemplate := templates.Layout(
		"Project Information", // define title text
		metaTags,              // define meta tags
		bodyContent,           // define body content
	)

	// Render info page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, infoTemplate); err != nil {
		// If not, return HTTP 400 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, return HTTP 400 error.
		slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Write HTML content.
	w.Write([]byte("<p>🎉 Yes, <strong>htmx</strong> is ready to use! (<code>GET /api/hello-world</code>)</p>"))

	// Send htmx response.
	htmx.NewResponse().Write(w)

	// Send log message.
	slog.Info("request API", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}
