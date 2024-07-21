// Custom terminal output for the application
package msg

import "fmt"

const Prompt = "\033[5m|> \033[0m"

// Custom errors for the application
type PokeError struct {
	Msg string
}

func (pe *PokeError) Error() string {
	return fmt.Sprintf("\033[31m%s\033[0m", pe.Msg)
}

func CmdNotFoundErr() PokeError {
	return PokeError{
		Msg: "Hmm... I am not sure what you mean. Try `help` if you are feeling lost.",
	}
}

// Custom printing for the help menu
func PrintCmdUsage(s string) {
	fmt.Printf("\t\x1b[32m%s\x1b[0m\n", s)
}

func PrintCmdName(s string) {
	fmt.Printf("\x1b[36m%s\x1b[0m\n", s)
}
