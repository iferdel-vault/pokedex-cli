package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "location-area/"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
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

	return LocationAreas, nil
}

func (l LocationAreasResp) ParseLocationNames() ([]byte, error) {
	fmt.Println("Location areas:")
	locationNames := ""
	for _, area := range l.Results {
		fmt.Printf("- %v\n", area.Name)
		locationNames += area.Name + "\n"
	}
	locationNames = locationNames[:len(locationNames)-1]
	return []byte(locationNames), nil
}
