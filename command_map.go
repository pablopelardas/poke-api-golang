package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config, _ []string) error {
	locationsResp, err := config.apiClient.ListLocations(config.NextPage)
	if err != nil {
		return err
	}
	config.NextPage = locationsResp.Next
	config.PrevPage = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}


func commandMapB(config *Config, _ []string) error {
	if config.PrevPage == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.apiClient.ListLocations(config.PrevPage)
	if err != nil {
		return err
	}

	config.NextPage = locationResp.Next
	config.PrevPage = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}