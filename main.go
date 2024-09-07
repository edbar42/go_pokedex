package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/edbar42/go_pokedex/api"
	"github.com/edbar42/go_pokedex/ui"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := api.NewCache(5 * time.Minute)

	fmt.Print(ui.Prompt)

	for scanner.Scan() {
		input := scanner.Text()
		command, exists := Commands[input]

		if exists {
			err := command.Callback(cache)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := ui.CmdNotFoundErr()
			fmt.Println(err.Error())
		}

		fmt.Print(ui.Prompt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
}
