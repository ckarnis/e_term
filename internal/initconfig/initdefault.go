package initconfig

type ConfigDefaults struct {
	Fred   APIConfig `toml:"fred"`
	Census APIConfig `toml:"census"`
}

type APIConfig struct {
	BaseURL        string     `toml:"base_url"`
	APIKey         string     `toml:"api_key"`
	TimeoutSeconds int        `toml:"timeout_seconds"`
	Endpoints      []Endpoint `toml:"endpoint"`
}

type Endpoint struct {
	Name string `toml:"name"`
	Path string `toml:"path"`
}

func defaultConfig() ConfigDefaults {
	return ConfigDefaults{
		Fred: APIConfig{
			BaseURL:        "https://api.stlouisfed.org/fred/",
			APIKey:         "",
			TimeoutSeconds: 10,
			Endpoints: []Endpoint{
				{Name: "series", Path: "/series"},
				{Name: "observations", Path: "/series/observations"},
			},
		},
		Census: APIConfig{
			BaseURL:        "",
			APIKey:         "",
			TimeoutSeconds: 10,
			Endpoints: []Endpoint{
				{Name: "population", Path: "/data/population"},
			},
		},
	}
}
