package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)


func  callbackCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon provided")
	}
	pokemonName := args[0]
	pokeBall := args[len(args) -1]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)

	if err != nil {
		fmt.Println(err)
	}
	threshold := 0
	switch  {
	case pokeBall == "basic":
		threshold = 50
	case pokeBall == "great":
		threshold = 80
	case pokeBall == "ultra":
		threshold = 200
	}
	caught := rand.IntN(pokemon.BaseExperience)
	caughtResult:=""
	if caught > threshold {
		caughtResult = "escaped"
	}
	if caught < threshold {
		caughtResult = "catched"
		cfg.caughtPokemon[pokemonName] = pokemon
	
	}
	fmt.Printf("Throwing a %v Pokeball at %v...\n",pokeBall,pokemonName)
	fmt.Printf("%v %v\n",pokemonName,caughtResult)
	
	return nil
}