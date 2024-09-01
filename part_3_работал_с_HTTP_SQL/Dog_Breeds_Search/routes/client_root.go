package routes

import (
	pages "dog_breeds_search/client/pages/home"
	"dog_breeds_search/utils"
	"net/http"
)

func ClientRoot(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			utils.ParseHtml(w, pages.MainHtml)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
