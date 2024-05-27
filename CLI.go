// CLI.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// config as a pointer in the paramater porque queremos shared access to its values
func CLI(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex >")
		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Command not available, see 'help'")
			continue
		}
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(input string) []string {
	words := strings.Fields(input)
	return words
}

// cliCommand is used to construct the CliCommands
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "explore {location-area}",
			description: "displays the previous 20 location areas in Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "catches a pokemon with a success rate based on pokemon base experience",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "inspect pokemon information if in Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all catched pokemons",
			callback:    commandPokedex,
		},
	}
}
