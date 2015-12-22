package static

import (
	"html/template"
	"path"

	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type staticRouter struct {
	cachedTemplates *template.Template
	routes          []router.Route
}

func NewRouter(templatesPath string) router.Router {
	r := &staticRouter{}

	r.initRoutes()
	r.cacheTemplates(templatesPath)

	return r
}

func (r *staticRouter) Routes() []router.Route {
	return r.routes
}

func (r *staticRouter) initRoutes() {
	r.routes = []router.Route{
		httputils.NewGetRoute(
			"/",
			r.Index,
			false,
			[]string{},
			httputils.Static,
		),
	}
}

func (r *staticRouter) cacheTemplates(templatesPath string) {
	index := path.Join(templatesPath, "/index.html")

	r.cachedTemplates = template.Must(template.ParseFiles(index))
}
