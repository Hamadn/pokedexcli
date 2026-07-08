package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func commandCatch(c *config, name string) error {
	if name == "" {
		return fmt.Errorf("Please provide a pokemon to catch")
	}
	fmt.Println("Throwing a Pokeball at " + name + "...")

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, err := fetchJSON(url)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	if rand.Intn(pokemon.BaseExperience+100) < 100 {
		fmt.Println(name + " was caught!")
		c.Pokedex[pokemon.Name] = pokemon
		fmt.Println(c.Pokedex)
	} else {
		fmt.Println(name + " escaped...")
	}

	return nil
}
