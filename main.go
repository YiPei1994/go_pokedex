package main

import (
	"time"

	"github.com/YiPei1994/PokeGo/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationURL  *string
	caughtPokemon map[string]pokeapi.PokemonDetail
}


func main() {
	cacheInterval := time.Minute * 5
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.PokemonDetail),
	}

	useRepl(&cfg)
}
