package routes

import (
	"assignment/sharable"
	"assignment/utils"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func GetBreeds(mux *http.ServeMux) {
	mux.HandleFunc("/getBreeds", func(w http.ResponseWriter, r *http.Request) {
		if len(*sharable.Dogs) == 0 {
			if r.Method == http.MethodGet {
				url := sharable.ApiUrl
				log.Println("url: ", url)
				client, err := http.DefaultClient.Get(url + "/breeds")
				if err != nil {
					utils.ResponseErrorText(err, w, "Error while getting breeds")
					return
				}
				defer client.Body.Close()
				body, readAllErr := io.ReadAll(client.Body)
				if readAllErr != nil {
					utils.ResponseErrorText(err, w, "Error while getting breeds from body")
					return
				}
				err = json.Unmarshal(body, &sharable.Dogs)
				if err != nil {
					utils.ResponseErrorText(err, w, "Failed to unmarshal breeds")
					return
				}

				for _, breed := range *sharable.Dogs {
					*sharable.DogBreeds = append(*sharable.DogBreeds, breed.Name)
				}
				log.Println("all bread: \n", sharable.DogBreeds)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			for _, breed := range *sharable.Dogs {
				*sharable.DogBreeds = append(*sharable.DogBreeds, breed.Name)
			}
		}
		if sharable.DogBreeds != nil {
			response, err := json.Marshal(sharable.DogBreeds)
			if err != nil {
				utils.ResponseErrorText(err, w, "Failed to marshal breeds")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		} else {
			utils.ResponseErrorText(errors.New("no breeds"), w, "No breeds")
		}
	})
}
