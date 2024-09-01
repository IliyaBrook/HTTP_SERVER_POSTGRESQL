package routes

import (
	"main/internal/handlers"
	"main/internal/middlewares"
	"net/http"
)

func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleProducts)))
}
