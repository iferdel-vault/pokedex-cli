package main

import (
	"bufio"
	"fmt"
	"os"
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

func mapCommand(p *Pagination) {

}

func mapbCommand(p *Pagination) {

}

type Pagination struct {
	page        int
	nextUrl     string
	previousUrl string
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
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in Pokemon world",
			callback:    mapbCommand,
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
