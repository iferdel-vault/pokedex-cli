package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAPI(endpoint string, jsonStructure interface{}) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error encountered during the http.Get method")
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("status code of get request:", res.StatusCode)
	if err != nil {
		fmt.Println("Error encountered during the read of the body from the response")
		log.Fatal(err)
	}

	err = json.Unmarshal(body, jsonStructure)
	if err != nil {
		fmt.Println(err)
	}

}
