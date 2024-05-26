package main

import (
	"errors"
	"fmt"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
	"github.com/iferdel/pokedexcli/internal/pokecache"
)

func commandMapf(c *config, cache *pokecache.Cache) {
	// check if data in cache already
	if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
		cache.Get(*c.currentEndPoint)
		fmt.Println(string(cachedData))
		return
	}
	pokeapi.GetAPI(*c.currentEndPoint, &c.locationAreas)
	locationValues := c.locationAreas.GetLocationNames()
	cache.Add(*c.currentEndPoint, []byte(locationValues))

	c.currentEndPoint = c.locationAreas.Next
}

func commandMapb(c *config, cache *pokecache.Cache) error {

	if c.locationAreas.Previous == nil {
		fmt.Println("you are in the first page")
		return errors.New("currently on first page")
	}

	c.currentEndPoint = c.locationAreas.Previous

	if cachedData, ok := cache.Get(*c.currentEndPoint); ok {
		cache.Get(*c.currentEndPoint)
		fmt.Println(string(cachedData))
		return nil
	}

	pokeapi.GetAPI(*c.currentEndPoint, &c.locationAreas)
	locationValues := c.locationAreas.GetLocationNames()
	cache.Add(*c.currentEndPoint, []byte(locationValues))

	return nil
}
