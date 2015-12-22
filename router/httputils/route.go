package httputils

type route struct {
	method        string
	path          string
	handler       APIHandler
	requiresAuth  bool
	requiredRoles []string
	routeType     RouteType
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Handler() APIHandler {
	return r.handler
}

func (r *route) RequiresAuth() bool {
	return r.requiresAuth
}

func (r *route) RequiredRoles() []string {
	return r.requiredRoles
}

func (r *route) Type() RouteType {
	return r.routeType
}

func NewRoute(method, path string, handler APIHandler, requiresAuth bool, requiredRoles []string, routeType RouteType) *route {
	return &route{
		method:        method,
		path:          path,
		handler:       handler,
		requiresAuth:  requiresAuth,
		requiredRoles: requiredRoles,
		routeType:     routeType,
	}
}

func NewGetRoute(path string, handler APIHandler, requiresAuth bool, requiredRoles []string, routeType RouteType) *route {
	return NewRoute("GET", path, handler, requiresAuth, requiredRoles, routeType)
}

func NewPostRoute(path string, handler APIHandler, requiresAuth bool, requiredRoles []string, routeType RouteType) *route {
	return NewRoute("POST", path, handler, requiresAuth, requiredRoles, routeType)
}
