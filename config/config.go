package config

import (
	"encoding/json"
)

// GTWConfiguration contains the API key for google maps
type GTWConfiguration struct {
	APIKey string `json:"apiKey"`
}

// Load the configuration from file
func Load(file []byte) (config *GTWConfiguration, err error) {
	err = json.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
