package initconfig

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func confirmOverwrite() (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("config.toml already exists. Overwrite? (y/N): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes", nil
}
