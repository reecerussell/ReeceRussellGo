package Authentication

import (
	"net/http"
	"strings"
)

// HandlerFunc ... http function in a simpler form
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Middleware ... authenticates http requests
func Middleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Pass straig through auth for development
		next(w, r)
		return

		header := r.Header.Get("Authorization")

		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		args := strings.Split(header, " ")
		scheme := args[0]
		token := args[1]

		if scheme != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		service := Service{}
		service.Init()

		success, err := service.ValidateToken(token)
		if err != nil || !success {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
