package routes

import (
	"main/internal/handlers"
	"main/internal/middlewares"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUsers)))
	mux.HandleFunc("/userProducts", middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUserProducts)))
}
