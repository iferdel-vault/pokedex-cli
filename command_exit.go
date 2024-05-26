package main

import (
	"os"
)

func commandExit(cfg *config, locationArea *string) (err error) {
	os.Exit(0)
	return nil
}
