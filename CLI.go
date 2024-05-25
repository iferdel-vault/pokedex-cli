package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/iferdel/pokedexcli/internal/pokeapi"
	"github.com/iferdel/pokedexcli/internal/pokecache"
)

type config struct {
	currentEndPoint *string
	locationAreas   *pokeapi.LocationAreasResp
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
			callback:    func() error { commandMapb(c, cache); return nil },
		},
	}
}

func cleanInput(input string) []string {
	words := strings.Fields(input)
	return words
}

func CLI() {
	scanner := bufio.NewScanner(os.Stdin)

	var initialLocationAreaEndpoint string = "https://pokeapi.co/api/v2/location-area/"
	config := &config{
		currentEndPoint: &initialLocationAreaEndpoint,
		locationAreas:   &pokeapi.LocationAreasResp{},
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
