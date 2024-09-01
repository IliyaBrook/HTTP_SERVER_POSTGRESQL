package routes

import (
	"dog_breeds_search/sharable"
	"dog_breeds_search/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// GetDogs
// @Summary Get dogs by breed
// @Description Returns a list of dogs by breed name. If the breed name is partially matched, similar results are returned.
// @Tags dogs
// @accept json
// @Produce  json
// @Param breed query string true "Breed name" default(Airedale Terrier)
// @Param x-api-key header string true "API key required to access this endpoint" default(sharable.ApiKey)
// @Success 200 {array} sharable.DogResponse "List of matching dogs"
// @Failure 400 {string} string "Breed name is required"
// @Failure 500 {string} string "Internal server error"
// @Router /getDogs [get]
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
