package main

import (
	"errors"
)

func commandMapf(cfg *config) error {
	/*
	   // check if data in cache already
	   if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
	       cache.Get(*c.currentEndPoint)
	       fmt.Println(string(cachedData))
	       return nil
	   }
	*/
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil

	//    cache.Add(*c.currentEndPoint, []byte(locationValues))

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
		return err
	}
	_, err = resp.ParseLocationNames()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	//cache.Add(*c.currentEndPoint, []byte(locationValues))

	return nil
}
