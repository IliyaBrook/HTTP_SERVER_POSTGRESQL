package sharable

type ImagesResponse struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DogResponse struct {
	Weight struct {
		Imperial string `json:"imperial"`
		Metric   string `json:"metric"`
	} `json:"weight"`
	Height struct {
		Imperial string `json:"imperial"`
		Metric   string `json:"metric"`
	} `json:"height"`
	Id               int    `json:"id"`
	Name             string `json:"name"`
	BredFor          string `json:"bred_for,omitempty"`
	BreedGroup       string `json:"breed_group,omitempty"`
	LifeSpan         string `json:"life_span"`
	Temperament      string `json:"temperament,omitempty"`
	Origin           string `json:"origin,omitempty"`
	ReferenceImageId string `json:"reference_image_id"`
	Image            struct {
		Id     string `json:"id"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Url    string `json:"url"`
	} `json:"image"`
	CountryCode string `json:"country_code,omitempty"`
	Description string `json:"description,omitempty"`
	History     string `json:"history,omitempty"`
}

var Dogs = &[]DogResponse{}
var DogBreeds = &[]string{}

type DogImageResponse struct {
	Breeds []struct {
		Weight struct {
			Imperial string `json:"imperial"`
			Metric   string `json:"metric"`
		} `json:"weight"`
		Height struct {
			Imperial string `json:"imperial"`
			Metric   string `json:"metric"`
		} `json:"height"`
		Id               int    `json:"id"`
		Name             string `json:"name"`
		CountryCode      string `json:"country_code"`
		BredFor          string `json:"bred_for"`
		BreedGroup       string `json:"breed_group"`
		LifeSpan         string `json:"life_span"`
		Temperament      string `json:"temperament"`
		Origin           string `json:"origin"`
		ReferenceImageId string `json:"reference_image_id"`
	} `json:"breeds"`
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
