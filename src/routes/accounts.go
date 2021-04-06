package routes

import (
	"challenge-golang-stone/src/controllers"
	"net/http"
)

var accountRoutes = []Route{
	{
		URI:        "/accounts",
		Method:     http.MethodGet,
		Function:   controllers.GetAccounts,
		AuthNeeded: false,
	},
	{
		URI:        "/accounts/{account_id}/balance",
		Method:     http.MethodGet,
		Function:   controllers.GetAccountBalance,
		AuthNeeded: false,
	},
	{
		URI:        "/accounts",
		Method:     http.MethodPost,
		Function:   controllers.CreateAccount,
		AuthNeeded: false,
	},
}
