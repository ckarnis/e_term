package commands

import (
	"ecoTerm/internal/config"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

func View(args []string) string {
	if len(args) == 0 {
		return "usage: view <source>"
	}

	var fileData map[string]any

	path := config.GetConfigPath("ecoTerm")

	if _, err := toml.DecodeFile(path, &fileData); err != nil {
		fmt.Println("can't read config.toml")
		os.Exit(1)
	}

	item, ok := fileData[args[0]]

	if !ok {
		return fmt.Sprintf("source %q not found", args[0])
	}

	m := item.(map[string]any)
	baseURL := m["base_url"].(string)
	apiKey := m["api_key"].(string)

	url := baseURL + "releases?api_key=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "error"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response: %v", err)
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
	return url
}
