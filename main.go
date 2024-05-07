package main

import (
	"time"

	"github.com/YiPei1994/PokeGo/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationURL  *string
}

func main() {
	cacheInterval := time.Minute * 5
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
	}

	useRepl(&cfg)
}
