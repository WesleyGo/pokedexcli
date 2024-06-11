package main

import (
	"time"

	"github.com/WESLEYGO/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		stats:         make(map[string]pokeapi.PokemonStats),
	}
	startRepl(cfg)
}
