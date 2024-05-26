package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) (err error) {
	fmt.Println("This is the help of the pokedex")
	fmt.Println("All the available commands are listed bellow:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
