package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Callback    func() error
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help": {
			Name:        "Help",
			Usage:       "help",
			Description: "Prints this help menu.",
			Callback:    help,
		},
		"exit": {
			Name:        "Exit",
			Usage:       "exit",
			Description: "Quits the program.",
			Callback:    exit,
		},
	}
}

func help() error {
	for _, cmd := range Commands {
		fmt.Println(cmd.Name)
		fmt.Println("\t", cmd.Usage)
		fmt.Println("\t", cmd.Description)
	}
	return nil
}

func exit() error {
	fmt.Println("Good luck catching them all. Bye.")
	os.Exit(0)
	return nil
}
