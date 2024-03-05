package main

import "os"

func commandExit(config *Config, _ []string) error {
	os.Exit(0)
	return nil
}