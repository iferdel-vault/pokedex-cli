// main.go
package main

import (
	"time"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
	"github.com/peterh/liner"
)

// statefull information for command callback functions
type config struct {
	CLILiner            *liner.State
	pokeapiClient       pokeapi.Client //to reuse the http client (more eff than constant creation)
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	Pokedex             pokeapi.Pokedex
}

func main() {
	cfg := &config{
		pokeapiClient:       pokeapi.NewClient(time.Minute, time.Hour), // rather than creating a new client for every new callback command
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
		Pokedex:             pokeapi.Pokedex{},
	}
	CLI(cfg)
}
