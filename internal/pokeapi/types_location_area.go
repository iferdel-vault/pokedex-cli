package pokeapi

// in its own file because its so specific to the domain we want to cover
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // Next is nil in last page
	Previous *string `json:"previous"` // Previous is nil in first page
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
