package main

import (
	"errors"
	"fmt"
)


func  callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		fmt.Println("no such pokemon in pokedex")
		return nil
	}
	fmt.Printf("Name: %v\n",pokemon.Name)
	fmt.Printf("Height: %v\n",pokemon.Height)
	fmt.Printf("Weight: %v\n",pokemon.Weight)
	fmt.Println("Stats: ")
	for _,v := range pokemon.Stats{
		fmt.Printf("	-%v: %v\n",v.Stat.Name,v.BaseStat)
	}
	fmt.Println("Type: ")
	for _,v := range pokemon.Types{
		fmt.Printf("	-%v\n",v.Type.Name)
	}
	
	
	return nil
}