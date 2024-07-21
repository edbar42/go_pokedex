package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/edbar42/go_pokedex/msg"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(msg.Prompt)

	for scanner.Scan() {
		input := scanner.Text()
		command, exists := Commands[input]

		if exists {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := msg.CmdNotFoundErr()
			fmt.Println(err.Error())
		}

		fmt.Print(msg.Prompt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
}
