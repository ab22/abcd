package handlers

import (
	"log"
	"net/http"
)

// Contains all handlers in charge of serving static pages and files.
type staticHandler struct{}

// Since Go's router sends all lost requests to home path '/',
// then we check if the URL path is not '/'.
// If the requested URL is '/', then we render the index.html template.
// If it's not, then we return a 404 response.
func (h *staticHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := cachedTemplates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "HTTP 500: Internal server error", http.StatusInternalServerError)
	}
}
