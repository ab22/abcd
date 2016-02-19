package static

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/httputils"
	"golang.org/x/net/context"
)

// Handler structure for the static handler.
type Handler struct {
	cachedTemplates *template.Template
}

// NewHandler creates a new router for the static handler.
func NewHandler(cfg *config.Config) *Handler {
	var (
		h     = &Handler{}
		index = path.Join(cfg.App.Frontend.Admin, "index.html")
	)

	h.cachedTemplates = template.Must(template.ParseFiles(index))

	return h
}

// Since Go's router sends all lost requests to home path '/',
// then we check if the URL path is not '/'.
// If the requested URL is '/', then we render the index.html template.
// If it's not, then we return a 404 response.
func (r *staticRouter) Index(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return nil
	}

	err := r.cachedTemplates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "HTTP 500: Internal server error", http.StatusInternalServerError)
	}

	return nil
}
