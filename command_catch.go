package main

import (
	"errors"
	"fmt"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon passed as argument to catch")
	}

	pokemon := args[0]

	resp, err := cfg.pokeapiClient.GetPokemonInfo(&pokemon)
	if err != nil {
		return err
	}

	pokedex := pokeapi.Pokedex{}
	pokedex[resp.Name] = resp

	fmt.Println("printing pokedex to date:", pokedex)

	return nil
}
