package commands

import (
	"eTerm/internal/config"
	"eTerm/internal/view"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func View(args []string) string {
	if len(args) == 0 {
		return "usage: view <source> <type> <dataset>"
	}

	var fileData map[string]any

	path := config.GetConfigPath("eTerm")

	if _, err := toml.DecodeFile(path, &fileData); err != nil {
		fmt.Println("can't read config.toml")
		os.Exit(1)
	}

	item, ok := fileData[args[0]]
	if !ok {
		return fmt.Sprintf("source %q not found", args[0])
	}

	//remove later
	println(args[0])
	return view.GetView(args, item)
}
