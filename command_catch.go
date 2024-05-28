package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	fmt.Printf("%s has %d base experience\n",
		resp.Name,
		resp.BaseExperience,
	)
	time.Sleep(1 * time.Second)

	const threshold = 50
	catchChance := rand.Intn(resp.BaseExperience)
	if catchChance > threshold {
		pokedex := cfg.Pokedex
		pokedex[resp.Name] = resp
		fmt.Printf("%s was caught!\n", resp.Name)
		fmt.Println("You may now inspect it with inspect command")
		return nil
	} else {
		return fmt.Errorf("%s escaped! ", resp.Name)
	}
}
