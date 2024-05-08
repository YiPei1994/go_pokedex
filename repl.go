package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func useRepl(cfg *config) []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type %v to learn how to work with program. \n","help")
	for {
		fmt.Print(">")

		scanner.Scan()

		input := scanner.Text()

		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		cName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		
		availableCommands := getComands()

		command, ok := availableCommands[cName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg,args...)
		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getComands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "get list of nearby locations for explore",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "get list of prev locations for explore",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "get list of pokemons in location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name} {pokeball_type}",
			description: "catching a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "inspect caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "check pokedex",
			callback:    callbackPokedex,
		},
		"choose": {
			name:        "choose {pokemon_name}",
			description: "choose pokemon for fight",
			callback:    callbackChoose,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
