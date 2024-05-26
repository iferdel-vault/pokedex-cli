package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

// Primer error a la vista, la función alude a un GET, pero no existe ningún return de esta función
// cuando por lo general un GET request tiene un return (los datos obtenidos del request)
func GetAPI(jsonStructure interface{}) {
	endpoint := InitialLocationAreaEndpoint
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
