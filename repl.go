package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func useRepl() []string {
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

		command.callback()

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getComands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name: "getMapLocations",
			description: "get next 20 map locations for explore",
			callback: callMapLocations,
		},
		"mapb":{

		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
