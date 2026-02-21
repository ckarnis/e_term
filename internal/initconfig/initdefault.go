package initconfig

type ConfigDefaults struct {
	Fred   APIConfig `toml:"fred"`
	Census APIConfig `toml:"census"`
}

type APIConfig struct {
	APIKey         string `toml:"api_key"`
	BaseURL        string `toml:"base_url"`
	TimeoutSeconds int    `toml:"timeout_seconds"`
}

func defaultConfig() ConfigDefaults {
	return ConfigDefaults{
		Fred: APIConfig{
			APIKey:         "",
			BaseURL:        "https://api.stlouisfed.org/fred",
			TimeoutSeconds: 10,
		},
		Census: APIConfig{
			APIKey:         "",
			BaseURL:        "",
			TimeoutSeconds: 10,
		},
	}
}
