package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(c *config, name string) error {
	fmt.Println("Throwing a Pokeball at " + name)
	if name == "" {
		return fmt.Errorf("Please provide a pokemon to catch")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, err := fetchJSON(url)
	if err != nil {
		return err
	}

	var pokemons Pokemon
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return err
	}

	if rand.Intn(pokemons.BaseExperience+100) < 100 {
		fmt.Println(name + " was caught!")
		c.Pokedex[pokemons.Name] = pokemons
	} else {
		fmt.Println(name + " escaped...")
	}

	return nil
}
