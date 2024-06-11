package main

import (
	"fmt"
)

func commandPokedex(conf *config) error {
	for _, pokemon := range conf.stats {
		fmt.Println("Your Pokedex:")
		fmt.Printf("Pokemon: %v\n", pokemon.Name)
	}
	return nil
}
