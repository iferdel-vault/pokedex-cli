package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemon *string) (Pokemon, error) {
	endpoint := fmt.Sprintf("pokemon/%s", *pokemon)
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making get request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error during 'Do' of request: %v", err)
	}
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("status code over 399: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return Pokemon{}, fmt.Errorf("error during reading of response body: %v", err)
	}

	PokemonInfo := Pokemon{}
	err = json.Unmarshal(body, &PokemonInfo)
	if err != nil {
		return PokemonInfo, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
	}

	return PokemonInfo, nil
}
