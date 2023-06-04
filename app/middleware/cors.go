package middleware

import (
	"net/http"
	"os"
)

func CORS(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedDomain := os.Getenv("ALLOWED_DOMAIN")

		w.Header().Set("Access-Control-Allow-Origin", allowedDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		handler(w, r)
	}
}
