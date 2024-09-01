package routes

import (
	"dog_breeds_search/sharable"
	"dog_breeds_search/utils"
	"encoding/json"
	"errors"
	"net/http"
)

// GetBreeds
// @Summary Get all breeds
// @Description Returns a list of all dog breeds.
// @Tags breeds
// @Produce json
// @Success 200 {array} string "List of all dog breeds"
// @Failure 500 {object} sharable.ErrorResponse "Internal server error"
// @Router /getBreeds [get]
func GetBreeds(mux *http.ServeMux) {
	mux.HandleFunc("/getBreeds", func(w http.ResponseWriter, r *http.Request) {
		if len(*sharable.Dogs) == 0 && len(*sharable.DogBreeds) == 0 {
			if r.Method == http.MethodGet {
				ok := utils.QueryGet(w, "/breeds", &sharable.Dogs)
				if !ok {
					return
				}
				for _, breed := range *sharable.Dogs {
					*sharable.DogBreeds = append(*sharable.DogBreeds, breed.Name)
				}
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			for _, breed := range *sharable.Dogs {
				*sharable.DogBreeds = append(*sharable.DogBreeds, breed.Name)
			}
		}
		if len(*sharable.DogBreeds) > 0 {
			response, err := json.Marshal(&sharable.DogBreeds)
			if err != nil {
				utils.ResponseErrorText(err, w, "Failed to marshal breeds")
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		} else {
			utils.ResponseErrorText(errors.New("no breeds"), w, "No breeds")
		}
	})
}
