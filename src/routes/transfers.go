package routes

import (
	"challenge-golang-stone/src/controllers"
	"net/http"
)

var transfersRoutes = []Route{
	{
		URI:        "/transfers",
		Method:     http.MethodGet,
		Function:   controllers.GetTransfers,
		AuthNeeded: true,
	},
	{
		URI:        "/transfers",
		Method:     http.MethodPost,
		Function:   controllers.Transfer,
		AuthNeeded: true,
	},
}
