package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadPokedex(cfg *config, filename string) error {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file")
	}
	err = json.Unmarshal(dat, &cfg.Pokedex)
	return nil
}

func savePokedex(cfg *config, filename string) error {
	data, err := json.MarshalIndent(cfg.Pokedex, "", "  ")
	if err != nil {
		return err
	}
	if f, err := os.Create(filename); err == nil {
		defer f.Close()
		_, err := f.Write(data)
		if err != nil {
			return fmt.Errorf("error saving data to pokedex before exiting CLI")
		}
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
