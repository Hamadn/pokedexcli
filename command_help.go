package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range commands() {
		fmt.Println(cmd.name, ":", cmd.desc)
	}
	return nil
}
