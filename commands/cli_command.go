package commands

import (
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
	orderedmap "github.com/wk8/go-ordered-map"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error
}

func CliCommands() *orderedmap.OrderedMap {
	m := orderedmap.New()
	m.Set(
		"help",
		&cliCommand{
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
	)
	m.Set(
		"exit",
		&cliCommand{
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    commandExit,
		},
	)
	m.Set(
		"map",
		&cliCommand{
			name:        "map",
			description: "Loads the next Pokemon areas page, if there is one",
			Callback:    commandMap,
		},
	)
	m.Set(
		"mapB",
		&cliCommand{
			name:        "mapB",
			description: "Loads the previous Pokemon areas page, if there is one",
			Callback:    commandMapB,
		},
	)
	m.Set(
		"explore",
		&cliCommand{
			name:        "explore",
			description: "Loads the list of Pokémon in a selected area",
			Callback:    commandExplore,
		},
	)
	m.Set(
		"catch",
		&cliCommand{
			name:        "catch",
			description: "Catches a Pokémon using its name and adds it to the Pokedex",
			Callback:    commandCatch,
		},
	)
	m.Set(
		"inspect",
		&cliCommand{
			name:        "inspect",
			description: "Lists the details of a previously caught Pokémon",
			Callback:    commandInspect,
		},
	)
	m.Set(
		"pokedex",
		&cliCommand{
			name:        "pokedex",
			description: "Lists all the names of the Pokémon the user has caught",
			Callback:    commandPokedex,
		},
	)
	return m
}

func Command(key string) (*cliCommand, bool) {
	command, found := CliCommands().Get(key)
	if found {
		return command.(*cliCommand), found
	}
	return nil, false
}
