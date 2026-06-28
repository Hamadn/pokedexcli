package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	desc     string
	callback func() error
}

func commands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:     "exit",
			desc:     "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name:     "help",
			desc:     "Prints help",
			callback: commandHelp,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range commands() {
		fmt.Println(cmd.name, ":", cmd.desc)
	}
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		switch cmd := words[0]; cmd {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		default:
			fmt.Println("Unknown command")
		}
	}
}
