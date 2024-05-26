package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResp, error) {

	endpoint := "location-area?offset=0&limit=20/"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// cache checking
	cachedValues, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("====cache hit=====")
		LocationAreas := LocationAreasResp{}
		err := json.Unmarshal(cachedValues, &LocationAreas)
		if err != nil {
			return LocationAreasResp{}, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
		}
		return LocationAreas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error making get request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error during 'Do' of request: %v", err)
	}
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("status code over 399: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error during reading of response body: %v", err)
	}

	LocationAreas := LocationAreasResp{}
	err = json.Unmarshal(body, &LocationAreas)
	if err != nil {
		return LocationAreas, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
	}

	c.cache.Add(fullURL, body)

	return LocationAreas, nil
}

func (c *Client) GetLocationArea(locationArea *string) (LocationAreaResp, error) {
	if locationArea == nil {
		return LocationAreaResp{}, fmt.Errorf("======")
	}

	endpoint := "location-area/" + *locationArea
	fullURL := baseURL + endpoint

	cachedValues, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("====cache hit=====")
		pokemonsInLocationArea := LocationAreaResp{}
		err := json.Unmarshal(cachedValues, &pokemonsInLocationArea)
		if err != nil {
			return pokemonsInLocationArea, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
		}

		return pokemonsInLocationArea, nil

	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("error making get request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("error during 'Do' of request: %v", err)
	}
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("status code over 399: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("error during reading of response body: %v", err)
	}

	pokemonsInLocationArea := LocationAreaResp{}
	err = json.Unmarshal(body, &pokemonsInLocationArea)
	if err != nil {
		return pokemonsInLocationArea, fmt.Errorf("error during unmarshal of body (JSON): %v", err)
	}

	c.cache.Add(fullURL, body)

	return pokemonsInLocationArea, nil

}
