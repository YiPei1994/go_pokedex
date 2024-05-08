package main

import (
	"fmt"
)


func  callbackPokedex(cfg *config, args ...string) error {
	for _,pokemon := range cfg.caughtPokemon {
		fmt.Println(pokemon.Name)
	}
	return nil
}