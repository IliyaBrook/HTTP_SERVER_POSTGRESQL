package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type assetsResponse struct {
	Data      []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type assetData struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

func main() {
	apiUrl := string("https://api.coincap.io/v2")
	limit := 5
	if //goland:noinspection ALL
	limit != 0 {
		apiUrl = apiUrl + "/assets?limit=" + strconv.Itoa(limit)
	} else {
		apiUrl = apiUrl + "/assets"
	}
	fmt.Println(apiUrl)
	client, err := http.DefaultClient.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(client.Body)

	if client.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code")
		fmt.Println("Status Code:", client.StatusCode)
		return
	}

	body, err := io.ReadAll(client.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Error: Failed to unmarshal JSON")
		fmt.Println("Error:", err)
		return
	}
	for _, asset := range r.Data {
		//fmt.Printf("%+v\n", asset)
		prettyJSON, _ := json.MarshalIndent(asset, "", "  ")
		fmt.Println(string(prettyJSON))
	}
}
