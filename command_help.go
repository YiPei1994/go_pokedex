package main

import "fmt"

func callbackHelp(cfg *config) error {
	availableCommands := getComands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range availableCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	fmt.Println("")

	return nil
}
