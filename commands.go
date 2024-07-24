package main

import (
	"fmt"
	"os"

	"github.com/edbar42/cfg"
	"github.com/edbar42/go_pokedex/msg"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Callback    func(cfg.CFG) error
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
		"map": {
			Name:        "Map",
			Usage:       "map",
			Description: "Lists the next page of mapped regions.",
			Callback:    mapn,
		},
	}
}

func help() error {
	for _, cmd := range Commands {
		fmt.Println("------------------")
		msg.PrintCmdName(cmd.Name)
		msg.PrintCmdUsage(cmd.Usage)
		fmt.Printf("\t%s\n", cmd.Description)
		fmt.Println("------------------")
	}
	return nil
}

func exit() error {
	fmt.Println("Good luck catching them all!")
	os.Exit(0)
	return nil
}

func mapn() error {
}
