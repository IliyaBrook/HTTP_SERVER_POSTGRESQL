package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// если нам не нужны кастомные параметры мы используем http.DefaultClient
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect")
			return nil
		},
	}
	//resp, err := http.DefaultClient.Get("https://jsonplaceholder.typicode.com/todos/1")
	resp, err := client.Get("https://www.baidu.com")

	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}
