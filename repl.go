package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/WESLEYGO/pokedexcli/internal/pokeapi"
)

func startRepl(conf *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		switch commandName {
		case "explore":
			if len(words) == 1 {
				fmt.Println("Please provide the area name")
				continue
			} else {
				conf.areaname = &words[1]
			}
		case "catch":
			if len(words) == 1 {
				fmt.Println("Please provide the pokemon")
				continue
			} else {
				conf.pokemon = &words[1]
			}
		case "inspect":
			if len(words) == 1 {
				fmt.Println("Please provide the pokemon")
				continue
			} else {
				conf.pokemon = &words[1]
			}
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	areaname      *string
	pokemon       *string
	stats         map[string]pokeapi.PokemonStats
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the locations in the Pokemon world",
			callback:    commandListLocations,
		},
		"mapb": {
			name:        "map",
			description: "List the previous locations in the Pokemon world",
			callback:    commandListPreviousLocations,
		},
		"explore": {
			name:        "explore",
			description: "List the Pokemon found the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon in the area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List a caught pokemon",
			callback:    commandPokedex,
		},
	}
}
