package middlewares

import (
	"context"
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

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-id")
		if userId == "" {
			log.Printf("[%s] %s - error: userID is not provided\n",
				r.Method, r.RequestURI,
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "ID", userId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
