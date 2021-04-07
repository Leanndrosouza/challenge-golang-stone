package middlewares

import (
	"challenge-golang-stone/src/auth"
	"challenge-golang-stone/src/responses"
	"net/http"
)

// Authenticate allow or deny access based on jwt token
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
