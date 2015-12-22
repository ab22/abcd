package httputils

type apiRoute struct {
	method        string
	path          string
	handler       APIHandler
	requiresAuth  bool
	requiredRoles []string
}

func (r *apiRoute) Path() string {
	return r.path
}

func (r *apiRoute) Method() string {
	return r.method
}

func (r *apiRoute) Handler() APIHandler {
	return r.handler
}

func (r *apiRoute) RequiresAuth() bool {
	return r.requiresAuth
}

func (r *apiRoute) RequiredRoles() []string {
	return r.requiredRoles
}

func NewRoute(method, path string, handler APIHandler, requiresAuth bool, requiredRoles []string) *apiRoute {
	return &apiRoute{
		method:        method,
		path:          path,
		handler:       handler,
		requiresAuth:  requiresAuth,
		requiredRoles: requiredRoles,
	}
}

func NewGetRoute(path string, handler APIHandler, requiresAuth bool, requiredRoles []string) *apiRoute {
	return NewRoute("GET", path, handler, requiresAuth, requiredRoles)
}

func NewPostRoute(path string, handler APIHandler, requiresAuth bool, requiredRoles []string) *apiRoute {
	return NewRoute("POST", path, handler, requiresAuth, requiredRoles)
}
