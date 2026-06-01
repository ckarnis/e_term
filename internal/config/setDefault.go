package config

type ConfigDefaults struct {
	Fred APIConfig `toml:"fred"`
}

type APIConfig struct {
	BaseURL        string     `toml:"base_url"`
	APIKey         string     `toml:"api_key"`
	TimeoutSeconds int        `toml:"timeout_seconds"`
	Endpoints      []Endpoint `toml:"endpoint"`
}

type RSSConfig struct {
	BaseURL   string     `toml:"base_url"`
	Endpoints []Endpoint `toml:"endpoint"`
}

type Endpoint struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
}

func DefaultConfig() ConfigDefaults {
	return ConfigDefaults{
		Fred: APIConfig{
			BaseURL: "https://api.stlouisfed.org/fred/",
			APIKey:  "",
			Endpoints: []Endpoint{
				{Name: "series", Path: "/series"},
				{Name: "observations", Path: "/series/observations"},
			},
		},
	}
}
