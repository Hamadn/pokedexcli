package main

import (
	"bufio"
	"fmt"
	"github.com/Hamadn/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(c *config, location string) error
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
		"map": {
			name:     "map",
			desc:     "List all maps",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Move to the previous map",
			callback: commandMapBack,
		},
		"explore": {
			name:     "explore",
			desc:     "Explore a location",
			callback: commandExplore,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Minute),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := commands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		location := ""
		if len(words) > 1 {
			location = words[1]
		}
		err := cmd.callback(cfg, location)
		if err != nil {
			fmt.Println(err)
		}
	}
}
