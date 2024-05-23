// main.go
package main

import (
	"fmt"

	"github.com/iferdel/pokedexcli/internal"
)

func main() {
	//CLI()

	const locationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	locationArea := internal.LocationArea{}

	internal.GetAPI(locationAreaEndpoint, &locationArea)

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
}
