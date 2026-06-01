package main

import (
	"log"
	"os"
	"strings"

	"ecoTerm/internal/app"
	"ecoTerm/internal/config"
	"ecoTerm/internal/windows"
)

func main() {

	if len(os.Args) >= 3 && os.Args[1] == "child" {

		name := os.Args[2]

		// starts with "config"
		if strings.HasPrefix(name, "config") {

			windows.RunConfigChild(name)

		} else {

			windows.RunData(name)
		}

		return
	}

	if err := config.Init("ecoTerm"); err != nil {
		log.Fatal(err)
	}

	app.Run()
}
