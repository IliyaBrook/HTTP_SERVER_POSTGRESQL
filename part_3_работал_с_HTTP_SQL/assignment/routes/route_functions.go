package routes

import (
	"net/http"
)

func RouteFunctions(mux *http.ServeMux) {
	ClientRoot(mux)
	GetBreeds(mux)
}
