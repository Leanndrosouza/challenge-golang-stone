package routes

import (
	"net/http"
)

// Route represents is a struct for all api routes
type Route struct {
	URI        string
	Method     string
	Function   func(http.ResponseWriter, *http.Request)
	AuthNeeded bool
}
