package main

import (
	"os"
)

func commandExit(cfg *config) (err error) {
	os.Exit(0)
	return nil
}
