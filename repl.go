package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	printTerminal()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			if err := processScannedText(scanner.Text()); err != nil {
				continue
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func processScannedText(text string) *emptyInputError {
	if len(text) == 0 {
		printTerminal()
		return &emptyInputError{}
	}
	cleanText := cleanedUpInput(text)
	if command, ok := cliCommands()[cleanText]; ok {
		if err := command.callback(); err != nil {
			fmt.Println(err)
			printTerminal()
		}
	} else {
		printUnknownCommand(cleanText)
		printTerminal()
	}
	return nil
}

func cleanedUpInput(input string) string {
	trimmed := strings.Trim(input, " ")
	return strings.ToLower(trimmed)
}

func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, command := range cliCommands() {
		fmt.Printf("%v: %v\n", name, command.description)
	}

	fmt.Println()

	printTerminal()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
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
