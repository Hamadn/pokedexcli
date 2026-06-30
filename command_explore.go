package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonEncounters struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, location string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	if location == "" {
		return fmt.Errorf("Please provide location to explore")
	}

	data, ok := c.cache.Get(url)
	var pokemon pokemonEncounters

	if ok {
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		jsonData, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		c.cache.Add(url, jsonData)

		err = json.Unmarshal(jsonData, &pokemon)
		if err != nil {
			return err
		}
	}

	fmt.Println("Found Pokemon:")
	for _, e := range pokemon.Encounters {
		fmt.Println("- " + e.Pokemon.Name)
	}

	return nil
}
