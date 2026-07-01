package main

import (
	"encoding/json"
	"fmt"
	"github.com/Hamadn/pokedexcli/internal/pokecache"
)

type config struct {
	nextUrl string
	prevUrl string
	cache   *pokecache.Cache
	Pokedex map[string]Pokemon
}

type locationArea struct {
	Next    string    `json:"next"`
	Prev    string    `json:"previous"`
	Results []mapItem `json:"results"`
}

type mapItem struct {
	Name string `json:"name"`
	Next string `json:"next"`
	Prev string `json:"previous"`
}

func commandMap(c *config, location string) error {
	url := c.nextUrl
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	data, ok := c.cache.Get(url)
	var err error

	if !ok {
		data, err = fetchJSON(url)
		if err != nil {
			return err
		}
		c.cache.Add(url, data)
	}
	var maps locationArea
	err = json.Unmarshal(data, &maps)
	if err != nil {
		return err
	}

	c.nextUrl = maps.Next
	c.prevUrl = maps.Prev

	for _, m := range maps.Results {
		fmt.Println(m.Name)
	}

	return nil
}

func commandMapBack(c *config, location string) error {
	if c.prevUrl == "" {
		return fmt.Errorf("you are on the first page")
	}
	c.nextUrl = c.prevUrl
	return commandMap(c, location)
}
