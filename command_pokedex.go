package main

import "fmt"

func commandPokedex(c *config, name string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
