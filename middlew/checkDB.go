package middlew

import (
	"github.com/Faiber-Barragan/twittor/db"
	"net/http"
)

//CheckDB is the middleware that allow to know state of the DB
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ConnectionCheck() == 0 {
			http.Error(w, "connection lost with the DB", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
