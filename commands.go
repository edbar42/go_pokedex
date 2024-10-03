package main

import (
	"fmt"
	"os"

	"github.com/edbar42/go_pokedex/api"
	"github.com/edbar42/go_pokedex/ui"
)

type Command struct {
	Name        string
	Usage       string
	Description string
	Callback    func(c *api.Cache, cmd_args []string) error
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
		"mapb": {
			Name:        "Map back",
			Usage:       "mapb",
			Description: "Lists the previous page of mapped regions.",
			Callback:    mapp,
		},
		"explore": {
			Name:        "Explore",
			Usage:       "explore [AREA_NAME]",
			Description: "Explore an area for Pokémons. To see area names, try map or mapb. ",
			Callback:    explore,
		},
	}
}

func help(c *api.Cache, cmd_args []string) error {
	for _, cmd := range Commands {
		fmt.Println("------------------")
		ui.PrintCmdName(cmd.Name)
		ui.PrintCmdUsage(cmd.Usage)
		fmt.Printf("\t%s\n", cmd.Description)
		fmt.Println("------------------")
	}

	fmt.Println(cmd_args)
	return nil
}

func exit(c *api.Cache, cmd_args []string) error {
	fmt.Println("Good luck catching them all!")
	os.Exit(0)
	return nil
}

func mapn(c *api.Cache, cmd_args []string) error {
	var reqURL string
	if c.NextMap == nil {
		reqURL = "https://pokeapi.co/api/v2/location/"
	} else {
		reqURL = *c.NextMap
	}

	return mapHelper(c, reqURL)
}

func mapp(c *api.Cache, cmd_args []string) error {
	if c.PrevMap == nil {
		noPrevErr := ui.NoPrevMapError()
		return &noPrevErr
	}

	return mapHelper(c, *c.PrevMap)
}

func explore(c *api.Cache, cmd_args []string) error {
	return nil
}
