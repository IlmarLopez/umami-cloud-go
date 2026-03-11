package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultEnv = "dev"

type EnvConfig struct {
	VPC VPC `json:"vps"`
}

type VPC struct {
	MaxAZs float64 `json:"max_azs"`
}

func LoadConfig() (*EnvConfig, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = defaultEnv
	}

	filePath := fmt.Sprintf("cdk/config/%s.json", env)

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s %w", filePath, err)
	}
	var config EnvConfig
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %w", err)
	}
	return &config, nil
}
