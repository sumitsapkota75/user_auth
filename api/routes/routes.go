package routes

import (
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(UserRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route Interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(userRoutes UserRoutes) Routes {
	return Routes{userRoutes}
}

// Setup all the route

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
