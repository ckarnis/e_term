package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {

	printBanner()

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("\033[32m❯❯\033[0m ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("bye")
			return
		}

		output := execute(input)

		if output != "" {
			fmt.Println("\033[90m" + output + "\033[0m")
		}
	}
}

func printBanner() {

	fmt.Println()
	fmt.Println("\033[32m[eco]Terminal\033[0m")
	fmt.Println("type 'help' for commands")
	fmt.Println()
}
