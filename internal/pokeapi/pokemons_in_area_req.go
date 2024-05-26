package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonsInLocationArea(locationArea *string) (PokemonsResp, error) {
	if locationArea == nil {
		return PokemonsResp{}, fmt.Errorf("======")
	}

	endpoint := "location-area/" + *locationArea
	fullURL := baseURL + endpoint

	cachedValues, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("====cache hit=====")
		pokemonsInLocationArea := PokemonsResp{}
		err := json.Unmarshal(cachedValues, &pokemonsInLocationArea)
		if err != nil {
			return pokemonsInLocationArea, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
		}

		return pokemonsInLocationArea, nil

	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return PokemonsResp{}, fmt.Errorf("error making get request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonsResp{}, fmt.Errorf("error during 'Do' of request: %v", err)
	}
	if resp.StatusCode > 399 {
		return PokemonsResp{}, fmt.Errorf("status code over 399: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return PokemonsResp{}, fmt.Errorf("error during reading of response body: %v", err)
	}

	pokemonsInLocationArea := PokemonsResp{}
	err = json.Unmarshal(body, &pokemonsInLocationArea)
	if err != nil {
		return pokemonsInLocationArea, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
	}

	c.cache.Add(fullURL, body)

	return pokemonsInLocationArea, nil

}
