package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/iferdel/pokedexcli/internal"
)

type CliCommands map[string]CliCommand

// CliCommand is used to construct the CliCommands
type CliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	fmt.Println("This is the help of the pokedex")
}

func commandExit() {
	os.Exit(0)
}

func commandMap(locationAreaEndpoint string, location *internal.LocationArea) {
	internal.GetAPI(locationAreaEndpoint, &location)
	for _, location := range location.Results {
		fmt.Println(location.Name)
	}
}

func commandMapb(locationAreaEndpoint string, location *internal.LocationArea) error {
	internal.GetAPI(locationAreaEndpoint, &location)
	if location.Previous == nil {
		fmt.Println("first page")
		return errors.New("you are on the first page")
	}
	return nil
}

func (c CliCommands) GetCommands() CliCommands {
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
			callback:    func() { commandMap(locationAreaEndpoint, &internal.LocationArea{}) },
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in Pokemon world",
			callback:    func() { commandMapb(locationAreaEndpoint, &internal.LocationArea{}) },
		},
	}
}

func CLI() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		c := CliCommands{}
		command, ok := c.GetCommands()[text]
		if !ok {
			fmt.Printf("Command not available, see 'help'\n")
			continue
		}
		command.callback()
	}
}
