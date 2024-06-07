package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"internal/pokecache"
	"os"
	"strings"
)

func startRepl(config *pokeapi.Config, cache *pokecache.Cache) {
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

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *pokecache.Cache) error
}

func cliCommands() map[string]cliCommand {
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
			description: "Load the next Pokemon areas page, if there is one",
			callback:    commandMap,
		},
		"mapB": {
			name:        "mapB",
			description: "Load the previous Pokemon areas page, if there is one",
			callback:    commandMapB,
		},
	}
}

func processScannedText(text string, config *pokeapi.Config, cache *pokecache.Cache) *emptyInputError {
	if len(text) == 0 {
		printTerminal()
		return &emptyInputError{}
	}
	cleanText := cleanedUpInput(text)
	if command, ok := cliCommands()[cleanText]; ok {
		if err := command.callback(config, cache); err != nil {
			fmt.Println(err)
			printTerminal()
		}
	} else {
		printUnknownCommand(text)
	}
	return nil
}

func cleanedUpInput(input string) string {
	return strings.Trim(input, " ")
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
