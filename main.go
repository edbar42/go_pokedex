package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "\033[5m|> \033[0m"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)

	for scanner.Scan() {
		input := scanner.Text()
		command, exists := Commands[input]

		if exists {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command", input)
		}

		fmt.Print(prompt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
}
