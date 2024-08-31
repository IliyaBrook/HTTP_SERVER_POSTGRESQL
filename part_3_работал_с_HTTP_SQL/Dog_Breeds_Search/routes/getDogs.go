package routes

import (
	"assignment/sharable"
	"assignment/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func GetDogs(mux *http.ServeMux) {
	mux.HandleFunc("/getDogs", func(w http.ResponseWriter, r *http.Request) {
		var dogs []sharable.DogResponse
		queryBreed := r.URL.Query().Get("breed")
		if queryBreed == "" {
			noBreedErr := errors.New("no breed")
			utils.ResponseErrorText(noBreedErr, w, "breed is required")
			return
		}
		encodedBreed := url.QueryEscape(queryBreed)
		ok := utils.QueryGet(w, "/breeds/search?name="+encodedBreed, &dogs)
		if !ok {
			return
		}
		if len(dogs) == 0 {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("[]"))
			return
		}

		jsonData, err := json.Marshal(&dogs)
		if err != nil {
			utils.ResponseErrorText(err, w, "Error marshal dogs to json")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
}
