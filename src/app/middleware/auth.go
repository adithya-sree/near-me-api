package middleware

import (
	"nearme-api/src/app/handler"
	"net/http"
)

//AuthMiddleware gets the auth middle
func (m Middleware) AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		if username != m.config.AppUsername || password != m.config.AppPassword {
			handler.RespondError(w, http.StatusUnauthorized, "invalid username/password")
			return
		}
		f(w, r)
	}
}
