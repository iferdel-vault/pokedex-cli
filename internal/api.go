package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// dejará de ser constante supongo, ya que offset nos permitirá cambiar de página
const locationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	res, err := http.Get(locationAreaEndpoint)
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

	pagedLocationArea := LocationArea{}
	err = json.Unmarshal(body, &pagedLocationArea)
	if err != nil {
		fmt.Println(err)
	}

}
