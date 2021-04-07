package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents is a struct for all api routes
type Route struct {
	URI        string
	Method     string
	Function   func(http.ResponseWriter, *http.Request)
	AuthNeeded bool
}

// Setup insert all routes in main router
func Setup(r *mux.Router) *mux.Router {
	routes := accountRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
