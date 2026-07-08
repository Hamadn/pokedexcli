package main

import "fmt"

func commandInspect(c *config, name string) error {
	if name == "" {
		return fmt.Errorf("Please provide a pokemon to catch")
	}

	fmt.Println("Inspecting the Pokedex...")

	pokemon, ok := c.Pokedex[name]
	if !ok {
		fmt.Println("Pokemon not caught")
		return nil
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Type:")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}
	return nil
}
