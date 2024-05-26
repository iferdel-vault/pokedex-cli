// CLI.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// el config del CLI debe tener información del client que se utilizará
// config con currentEndpoint puede ser, pero no es según pauta. Revisar
// config con locationAreas no corresponde, porque es un coupling muy grande.
// config as a pointer in the paramater porque queremos shared access to its values
func CLI(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("pokedex >")
		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			fmt.Printf("")
			continue
		}

		command, ok := getCommands()[cleanedInput[0]]
		if !ok {
			fmt.Printf("Command not available, see 'help'\n")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	words := strings.Fields(input)
	return words
}

type CliCommands map[string]CliCommand

// CliCommand is used to construct the CliCommands
type CliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() CliCommands {
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
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in Pokemon world",
			callback:    commandMapb,
		},
	}
}
