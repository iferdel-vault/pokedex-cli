package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	/*
		// check if data in cache already
		if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
			cache.Get(*c.currentEndPoint)
			fmt.Println(string(cachedData))
			return
		}
	*/

	resp, err := cfg.pokeapiClient.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	_, err = resp.ParseLocationNames()

	return nil

	/*
	   pokeapi.GetAPI(&c.locationAreas)
	   locationValues := c.locationAreas.GetLocationNames()
	   cache.Add(*c.currentEndPoint, []byte(locationValues))

	   c.currentEndPoint = c.locationAreas.Next
	*/
}

func commandMapb(cfg *config) error {

	if cfg.locationAreas.Previous == nil {
		fmt.Println("you are in the first page")
		return errors.New("currently on first page")
	}

	cfg.currentEndPoint = cfg.locationAreas.Previous

	/*
		if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
			cache.Get(*c.currentEndPoint)
			fmt.Println(string(cachedData))
			return nil
		}
	*/

	pokeapi.GetAPI(&cfg.locationAreas)
	// locationValues, _ := cfg.locationAreas.ParseLocationNames()
	//cache.Add(*c.currentEndPoint, []byte(locationValues))

	return nil
}
