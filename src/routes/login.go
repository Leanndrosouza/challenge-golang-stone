package routes

import (
	"challenge-golang-stone/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:        "/login",
	Method:     http.MethodPost,
	Function:   controllers.Login,
	AuthNeeded: false,
}
