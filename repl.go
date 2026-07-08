package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Hamadn/pokedexcli/internal/pokecache"
	"github.com/peterh/liner"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(c *config, location string) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"catch": {
			name:     "catch",
			desc:     "Catch a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect",
			desc:     "Inspect a pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "pokedex",
			desc:     "List your pokedex",
			callback: commandPokedex,
		},
	}
}

var history_fn = filepath.Join(os.TempDir(), "pokedex_history")

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	cfg := &config{
		cache:   pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]Pokemon),
	}

	cmds := commands()

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for k := range cmds {
			if strings.HasPrefix(k, strings.ToLower(line)) {
				c = append(c, k)
			}
		}
		return
	})

	if file, err := os.Open(history_fn); err == nil {
		line.ReadHistory(file)
		file.Close()
	}

	for {
		input, err := line.Prompt("Pokedex> ")
		if err != nil {
			break
		}
		line.AppendHistory(input)
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmd, ok := cmds[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		location := ""
		if len(words) > 1 {
			location = words[1]
		}
		err = cmd.callback(cfg, location)

		if errors.Is(err, ErrExit) {
			break
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	if file, err := os.Create(history_fn); err == nil {
		line.WriteHistory(file)
		file.Close()
	}
}
