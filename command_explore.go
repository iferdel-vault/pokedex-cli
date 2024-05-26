package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationArea *string) error {

	resp, err := cfg.pokeapiClient.GetPokemonsInLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("exploring pokemons in %q\n", string(*locationArea))
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
