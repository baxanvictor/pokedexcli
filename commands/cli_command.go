package commands

import (
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error
}

func CliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Loads the next Pokemon areas page, if there is one",
			Callback:    commandMap,
		},
		"mapB": {
			name:        "mapB",
			description: "Loads the previous Pokemon areas page, if there is one",
			Callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Loads the list of Pokémon in a selected area",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a Pokémon using its name and adds it to the Pokedex",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists the details of a previously caught Pokémon",
			Callback:    commandInspect,
		},
	}
}
