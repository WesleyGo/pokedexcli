package main

import (
	"fmt"
)

func commandExplore(conf *config) error {
	locationsResp, err := conf.pokeapiClient.ExploreArea(conf.areaname)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Printf("- %v\n", loc.Pokemon.Name)
	}
	return nil
}
