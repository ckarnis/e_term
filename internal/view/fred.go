package view

import (
	"ecoTerm/internal/fred"
)

func ViewFred(args []string, item any) string {
	m := item.(map[string]any)
	baseURL := m["base_url"].(string)
	apiKey := m["api_key"].(string)
	if len(args) == 0 {
		return "missing type\n\nusage: view <source> <type> <dataset>"
	}

	switch args[0] {
	case "releases":
		return fred.ViewReleases(baseURL, apiKey)
	default:
		return "something is wrong"
	}

}
