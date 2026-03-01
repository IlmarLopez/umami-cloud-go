package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

const defaultEnv = "dev"

type EnvConfig struct {
	VpcName string  `json:"vpcName"`
	Cidr    string  `json:"cidr"`
	MaxAzs  float64 `json:"maxAzs"`
}

func loadConfig() (*EnvConfig, error) {
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

func main() {
	defer jsii.Close()

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)

	envName := os.Getenv("ENV")
	if envName == "" {
		envName = defaultEnv
	}
	stackName := fmt.Sprintf("%sNetworkStack", envName)

	NewNetworkStack(app, stackName, &NetworkStackProps{
		StackProps: awscdk.StackProps{Env: env()},
		Config:     config,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
