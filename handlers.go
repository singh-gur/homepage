package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/singh-gur/homepage/templates"
	"github.com/singh-gur/homepage/templates/components"
	"github.com/singh-gur/homepage/templates/pages"

	"github.com/gofiber/fiber/v2"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(c *fiber.Ctx) error {

	// Define template functions.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	navbar := components.Navbar([]string{"Home", "About", "Contact"}, "Home") // define navbar items
	index := pages.Index2Content()
	footer := components.Footer(map[string]string{
		"GitHub":   "https://github.com/singh-gur",
		"LinkedIn": "https://www.linkedin.com/in/gurbakhshish-singh-1b73b177",
	})
	// Define template handler.
	templateHandler := templ.Handler(
		templates.Layout(
			"Welcome to example!", // define title text
			metaTags, navbar, index, footer,
		),
	)

	// Render template layout.
	return adaptor.HTTPHandler(templateHandler)(c)

}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(c *fiber.Ctx) error {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if c.Get("HX-Request") == "" || c.Get("HX-Request") != "true" {
		// If not, return HTTP 400 error.
		return fiber.NewError(fiber.StatusBadRequest, "non-htmx request")
	}

	return c.SendString("<p>🎉 Yes, <strong>htmx</strong> is ready to use! (<code>GET /api/hello-world</code>)</p>")
}
