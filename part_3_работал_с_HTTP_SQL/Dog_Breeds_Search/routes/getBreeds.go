package routes

import (
	"assignment/sharable"
	"assignment/utils"
	"encoding/json"
	"errors"
	"net/http"
)

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
