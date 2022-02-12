package middleware

import (
	"github.com/Crack1/twittor/db"
	"net/http"
)

/* CheckDB is the middleware that allows knows DB status */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ConectionCheck() == 0 {
			http.Error(w, "DB connection Lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
