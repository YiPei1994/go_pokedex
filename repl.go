package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func useRepl(cfg *config) []string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()

		input := scanner.Text()

		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		cName := cleaned[0]
		availableCommands := getComands()

		command, ok := availableCommands[cName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
