package main

import (
	"os"
)

func commandExit(cfg *config, args ...string) (err error) {
	os.Exit(0)
	return nil
}
