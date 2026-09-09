package view

func GetView(args []string, item any) string {
	m := item.(map[string]any)
	baseURL := m["base_url"].(string)
	apiKey := m["api_key"].(string)
	if len(args) == 0 {
		return "missing type\n\nusage: view <source> <type> <dataset>"
	}

	query := args[1]
	// next request from api
	// return just a place holder
	return baseURL + apiKey + "/" + query
}
