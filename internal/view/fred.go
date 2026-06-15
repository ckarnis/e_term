package view

import (
	"ecoTerm/internal/fred"
	"fmt"
	"io"
	"net/http"
)

func ViewFred(args []string, item any) string {
	m := item.(map[string]any)
	baseURL := m["base_url"].(string)
	apiKey := m["api_key"].(string)
	if len(args) == 0 {
		return "missing type\n\nusage: view <source> <type> <dataset>"
	}
	url := baseURL + "releases?api_key=" + apiKey

	switch args[0] {
	case "releases":
		return fred.ViewReleases(baseURL, apiKey)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "error"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("error reading response: %v", err)
	}

	fmt.Println("status:", resp.Status)
	fmt.Println("response:", string(body))

	return url
}
