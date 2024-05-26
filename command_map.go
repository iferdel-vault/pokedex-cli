package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are currently in the first page")
	}

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}

	return nil
}
