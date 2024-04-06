package simpleBubble

import "fmt"

func clearTerminal() {
	fmt.Print("\033[2J\033[H") // ANSI escape codes for clearing the screen
}
