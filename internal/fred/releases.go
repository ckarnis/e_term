package fred

func ViewReleases(baseURL string, apiKey string) string {
	url := baseURL + "releases?api_key=" + apiKey

	return url
}
