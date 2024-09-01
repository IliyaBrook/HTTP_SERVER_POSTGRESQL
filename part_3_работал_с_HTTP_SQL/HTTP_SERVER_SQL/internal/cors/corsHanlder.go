package cors

import (
	"github.com/rs/cors"
	"net/http"
)

func CORSHandler(mux *http.ServeMux) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}).Handler

	return corsHandler(mux)
}
