package main

import (
	"fmt"
)


func  callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("no pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _,pokemon := range cfg.caughtPokemon {
		fmt.Println(pokemon.Name)
	}
	return nil
}