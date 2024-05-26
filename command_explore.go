package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationAreaName *string) error {

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("exploring pokemons in %q\n", string(*locationAreaName))
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
