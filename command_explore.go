package main

import (
	"errors"
	"fmt"
)

func  callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	res, err := cfg.pokeapiClient.LocationAreaDetail(args[0])

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Exploring %v...\n",args)
	fmt.Println("Found Pokemon:")
	for _,v := range res.PokemonEncounters {
		fmt.Printf("- %v\n",v.Pokemon.Name)
	}
	return nil
}