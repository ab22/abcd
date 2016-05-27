package static

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the Static Handler.
type Handler interface {
	Index(context.Context, http.ResponseWriter, *http.Request) error
}

// handler structure for the static handler.
type handler struct {
	cachedTemplates *template.Template
}

// NewHandler initializes a static handler struct.
func NewHandler(cfg *config.Config) Handler {
	var (
		h     = &handler{}
		index = path.Join(cfg.App.Frontend.Admin, "index.html")
	)

	h.cachedTemplates = template.Must(template.ParseFiles(index))

	return h
}

// Index simply returns the index.html file from the main app.
//
// Since Go's router sends all lost requests to home path '/',
// then we check if the URL path is not '/'.
// If the requested URL is '/', then we render the index.html template.
// If it's not, then we return a 404 response.
func (h handler) Index(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return nil
	}

	err := h.cachedTemplates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "HTTP 500: Internal server error", http.StatusInternalServerError)
	}

	return nil
}
