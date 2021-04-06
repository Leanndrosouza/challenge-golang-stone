package routes

import (
	"net/http"
)

var accountRoutes = []Route{
	{
		URI:    "/accounts",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthNeeded: false,
	},
	{
		URI:    "/accounts/{account_id}/balance",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthNeeded: false,
	},
	{
		URI:    "/accounts",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthNeeded: false,
	},
}
