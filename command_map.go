package main

import (
	"errors"
	"log"
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
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil

	/*
	   pokeapi.GetAPI(&c.locationAreas)
	   locationValues := c.locationAreas.GetLocationNames()
	   cache.Add(*c.currentEndPoint, []byte(locationValues))

	   c.currentEndPoint = c.locationAreas.Next
	*/
}

func commandMapb(cfg *config) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are currently in the first page")
	}

	/*
		if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
			cache.Get(*c.currentEndPoint)
			fmt.Println(string(cachedData))
			return nil
		}
	*/

	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	//cache.Add(*c.currentEndPoint, []byte(locationValues))

	return nil
}
