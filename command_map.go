package main

import (
	"fmt"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
	"github.com/iferdel/pokedexcli/internal/pokecache"
)

func commandMap(c *config, cache *pokecache.Cache) {
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
