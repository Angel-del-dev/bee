package agent

import "fmt"

func Logo() string {
	return `
██████╗ ███████╗███████╗
██╔══██╗██╔════╝██╔════╝
██████╦╝█████╗  █████╗
██╔══██╗██╔══╝  ██╔══╝
██████╦╝███████╗███████╗
╚═════╝ ╚══════╝╚══════╝
🐝 `
}

func Init() {
	fmt.Print("\033[H\033[2J") // Clear terminal
	fmt.Print("\033[33m")
	fmt.Print(Logo(), RandomMessage())
	fmt.Println()
	fmt.Print("\033[0m")
}
