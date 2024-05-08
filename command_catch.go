package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)


func  callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)

	if err != nil {
		fmt.Println(err)
	}
	const threshold = 50
	caught := rand.IntN(pokemon.BaseExperience)
	caughtResult:=""
	if caught > threshold {
		caughtResult = "escaped"
	}
	if caught < threshold {
		caughtResult = "catched"
		cfg.caughtPokemon[pokemonName] = pokemon
	
	}
	fmt.Printf("Throwing a Pokeball at %v...\n",pokemonName)
	fmt.Printf("%v %v\n",pokemonName,caughtResult)
	
	return nil
}