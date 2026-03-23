package query

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
)

type Provider struct {
	APIKey         string `toml:"api_key"`
	BaseURL        string `toml:"base_url"`
	TimeoutSeconds int    `toml:"timeout_seconds"`
}

type Config map[string]Provider

const configPath = "config.toml"

func QueryProvider(name string, endpoint string, numOfObs string) (map[string]interface{}, error) {
	nums, err := strconv.Atoi(numOfObs)
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	provider, ok := cfg[name]
	if !ok {
		return nil, fmt.Errorf("provider not found: %s", name)
	}

	if provider.BaseURL == "" {
		return nil, fmt.Errorf("base_url is empty for provider: %s", name)
	}

	client := &http.Client{
		Timeout: time.Duration(provider.TimeoutSeconds) * time.Second,
	}

	url := provider.BaseURL + endpoint

	switch name {
	case "fred":
		url += fmt.Sprintf(
			"&api_key=%s&file_type=json&sort_order=desc&limit=%d",
			provider.APIKey,
			nums,
		)
	case "census":
		url += "no"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return nil, err
	}

	return dat, nil
}
