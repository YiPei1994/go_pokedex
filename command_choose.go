package main

import (
	"errors"
	"fmt"
)


func  callbackChoose(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon name provide")
	}
	pokemonName := args[0]
	pokemon,ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you havent got this pokemon yet")
	}

	cfg.choosedPokemon = pokemon

	fmt.Printf("fight %v, i choose you!\n",cfg.choosedPokemon.Name)
	
	return nil
}