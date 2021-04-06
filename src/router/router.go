package router

import (
	"challenge-golang-stone/src/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router with configured routers
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
