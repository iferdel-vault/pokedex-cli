package main

import (
	"os"
)

func commandExit(cfg *config, args ...string) (err error) {
	cfg.CLILiner.Close()
	os.Exit(0)
	return nil
}
