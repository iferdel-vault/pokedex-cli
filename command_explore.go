package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area passed as argument for explore command")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("exploring pokemons in %q\n", string(resp.Name))
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
