package parsers

import (
	"encoding/json"
	"io/ioutil"
)

// WebsiteAssertion handles the config tests
type WebsiteAssertion struct {
	Status   int      `json:"status,omitempty"`
	Selector string   `json:"selector,omitempty"`
	Contains []string `json:"contains,omitempty"`
}

// WebsiteConfig handles the config
type WebsiteConfig struct {
	Website    string             `json:"website"`
	URL        string             `json:"url"`
	Assertions []WebsiteAssertion `json:"assertions"`
}

// ReadConfig returns the full configuration array containing all the informations to be tested
func ReadConfig(configPath string) ([]WebsiteConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return parseJSON(data)
}

func parseJSON(jsonData []byte) ([]WebsiteConfig, error) {
	var config []WebsiteConfig
	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
