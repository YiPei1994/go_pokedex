package main

import "github.com/YiPei1994/PokeGo/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationURL  *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	useRepl(&cfg)
}
