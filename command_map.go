package main

import (
	"errors"
)

func commandMapf(cfg *config) error {

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

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
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
