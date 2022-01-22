package cors

import (
	"log"
	"net/http"
)

func setupCORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	log.Println(req.Header.Get("Origin"))
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupCORS(w, r)
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
