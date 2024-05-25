package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/iferdel/pokedexcli/internal/api"
	"github.com/iferdel/pokedexcli/internal/pokecache"
)

type config struct {
	currentEndPoint string
	locationArea    *internal.LocationArea
}

func commandMap(c *config, cache *pokecache.Cache) {
	// check if data in cache already
	if cachedData, ok := cache.Get(c.currentEndPoint); ok {
		cache.Get(c.currentEndPoint)
		fmt.Println(string(cachedData))
		return
	}
	internal.GetAPI(c.currentEndPoint, c.locationArea)
	locationValues := c.locationArea.GetLocationNames()
	cache.Add(c.currentEndPoint, []byte(locationValues))

	c.currentEndPoint = c.locationArea.Next
}

func commandMapb(c *config) error {

	if c.locationArea.Previous == nil {
		fmt.Println("you are in the first page")
		return errors.New("currently on first page")
	}

	c.currentEndPoint = *c.locationArea.Previous
	internal.GetAPI(c.currentEndPoint, c.locationArea)
	c.locationArea.GetLocationNames()

	return nil
}

type CliCommands map[string]CliCommand

// CliCommand is used to construct the CliCommands
type CliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands(c *config, cache *pokecache.Cache) CliCommands {
	return CliCommands{
		"help": {
			name:        "help",
			description: "this is the help of the pokedex",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the next 20 location areas in Pokemon world",
			callback:    func() error { commandMap(c, cache); return nil },
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in Pokemon world",
			callback:    func() error { commandMapb(c); return nil },
		},
	}
}

func cleanInput(input string) []string {
	words := strings.Fields(input)
	return words
}

func CLI() {
	scanner := bufio.NewScanner(os.Stdin)

	var initialLocationAreaEndpoint string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	config := &config{
		currentEndPoint: initialLocationAreaEndpoint,
		locationArea:    &internal.LocationArea{},
	}

	// cache initialization
	cache := pokecache.NewCache(10 * time.Second)

	for {
		fmt.Printf("pokedex >")
		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			fmt.Printf("")
			continue
		}

		command, ok := getCommands(config, cache)[cleanedInput[0]]
		if !ok {
			fmt.Printf("Command not available, see 'help'\n")
			continue
		}
		command.callback()
	}
}
