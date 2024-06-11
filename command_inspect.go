package main

import (
	"fmt"
)

func commandInspect(conf *config) error {
	if pokemon, ok := conf.stats[*conf.pokemon]; ok {
		fmt.Printf("Pokemon: %v\n", pokemon.Name)
		fmt.Printf("Base experience: %v\n", pokemon.BaseExperience)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf("-%v \n", t.Type.Name)
		}
		for _, s := range pokemon.Stats {
			fmt.Printf("-%v: %v \n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println()
	} else {
		fmt.Println("Pokemon not found")
	}
	return nil
}
