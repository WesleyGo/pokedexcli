package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config) error {
	stats, err := conf.pokeapiClient.CatchPokemon(conf.pokemon)
	if err != nil {
		return err
	}
	probability := 50 * stats.BaseExperience / 100

	roll := rand.Intn(100)

	fmt.Println("Throwing a Pokeball at pikachu...")

	if roll > probability {
		fmt.Println("Pokemon caught!")

	} else {
		fmt.Println("Pokemon got away!")
	}

	fmt.Printf("Base experience: %v\n", stats.BaseExperience)
	fmt.Printf("Probability: %v\n Roll: %v\n", probability, roll)

	conf.stats[*conf.pokemon] = stats

	return nil
}
