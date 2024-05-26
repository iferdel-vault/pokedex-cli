package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAPI(endpoint string, jsonStructure interface{}) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error encountered during the http.Get method")
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

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
