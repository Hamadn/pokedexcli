package main

import (
	"errors"
	"fmt"
)

var ErrExit = errors.New("exit")

func commandExit(c *config, location string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return ErrExit
}
