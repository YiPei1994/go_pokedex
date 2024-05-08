package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config,args ...string) error {

	res, err := cfg.pokeapiClient.ListLocationsAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	locations := res.Results
	fmt.Println("Nearby locations:")
	for _, area := range locations {
		fmt.Printf("- %v\n", area.Name)
	}
	cfg.nextLocationsURL = res.Next
	cfg.prevLocationURL = res.Previous
	return nil
}

func callbackMapB(cfg *config,args ...string) error {
	if cfg.prevLocationURL == nil {

		return errors.New("you are on page 0")
	}
	res, err := cfg.pokeapiClient.ListLocationsAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	locations := res.Results
	fmt.Println("Nearby locations:")
	for _, area := range locations {
		fmt.Printf("- %v\n", area.Name)
	}
	cfg.prevLocationURL = res.Previous
	cfg.nextLocationsURL = res.Next
	return nil
}
