package routes

import (
	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/services"

	"github.com/ab22/abcd/handlers/auth"
	"github.com/ab22/abcd/handlers/cart"
	"github.com/ab22/abcd/handlers/city"
	"github.com/ab22/abcd/handlers/country"
	"github.com/ab22/abcd/handlers/currency"
	"github.com/ab22/abcd/handlers/event"
	"github.com/ab22/abcd/handlers/eventtype"
	"github.com/ab22/abcd/handlers/state"
	"github.com/ab22/abcd/handlers/static"
	"github.com/ab22/abcd/handlers/user"
)

// Routes contains all Template and API routes for the application.
type Routes struct {
	TemplateRoutes []Route
	APIRoutes      []Route
}

// NewRoutes creates a new Router instance and initializes all template
// and API Routes.
func NewRoutes(cfg *config.Config, services *services.Services) *Routes {
	var (
		staticHandler    = static.NewHandler(cfg)
		authHandler      = auth.NewHandler(services)
		cartHandler      = cart.NewHandler(services)
		cityHandler      = city.NewHandler(services)
		countryHandler   = country.NewHandler(services)
		currencyHandler  = currency.NewHandler(services)
		eventHandler     = event.NewHandler(services)
		eventTypeHandler = eventtype.NewHandler(services)
		stateHandler     = state.NewHandler(services)
		userHandler      = user.NewHandler(services)

		r = &Routes{
			TemplateRoutes: []Route{},
			APIRoutes:      []Route{},
		}
	)

	return r
}
