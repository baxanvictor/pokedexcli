package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vbaxan-linkedin/pokedexcli/commands"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func startRepl(config *pokeapi.Config, cache *pokecache.AppCache) {
	printTerminal()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			if err := processScannedText(scanner.Text(), config, cache); err != nil {
				continue
			} else {
				printTerminal()
			}
		}
	}
}

func processScannedText(text string, config *pokeapi.Config, cache *pokecache.AppCache) error {
	if len(text) == 0 {
		printTerminal()
		return &emptyInputError{}
	}
	c, args := cleanedUpInput(text)
	if command, ok := commands.Command(c); ok {
		if err := command.Callback(config, cache, args...); err != nil {
			fmt.Println(err)
		}
	} else {
		printUnknownCommand(c)
	}
	return nil
}

func cleanedUpInput(input string) (command string, args []string) {
	fields := strings.Fields(input)
	switch fieldsLen := len(fields); fieldsLen {
	case 0:
		return
	case 1:
		return fields[0], nil
	default:
		return fields[0], fields[1:]
	}
}

func printTerminal() {
	fmt.Print("Pokedex > ")
}

func printUnknownCommand(command string) {
	fmt.Printf("Unknown command: %s\n", command)
}

type emptyInputError struct{}

func (eie *emptyInputError) Error() string {
	return "Empty input"
}
