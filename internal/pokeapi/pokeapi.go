package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string  `json:"next"` // Next is nil in last page
	Previous *string `json:"previous"` // Previous is nil in first page
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (l LocationAreasResp) GetLocationNames() []byte {
	locationNames := ""
	for _, location := range l.Results {
		fmt.Println(location.Name)
		locationNames += location.Name + "\n"
	}
	locationNames = locationNames[:len(locationNames)-1]
	return []byte(locationNames)
}

func GetAPI(endpoint string, jsonStructure interface{}) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error encountered during the http.Get method")
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Errorf("response comes with status code %q, expected 200", res.StatusCode)
	}
	if err != nil {
		fmt.Println("Error encountered during the read of the body from the response")
		log.Fatal(err)
	}

	err = json.Unmarshal(body, jsonStructure)
	if err != nil {
		fmt.Println(err)
	}

}
