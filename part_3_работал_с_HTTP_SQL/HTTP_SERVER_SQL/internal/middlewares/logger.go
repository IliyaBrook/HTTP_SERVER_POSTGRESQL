package middlewares

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("ID")
		userId, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userID is invalid", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("[%s] %s by userId %s\n", r.Method, r.URL, userId)
		next(w, r)
	}
}
