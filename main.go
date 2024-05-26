// main.go
package main

import (
	"time"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
)

// statefull information for command callback functions
type config struct {
	pokeapiClient       pokeapi.Client //to reuse the http client (more eff than constant creation)
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Minute, time.Hour), // rather than creating a new client for every new callback command
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}
	CLI(&cfg)
}
