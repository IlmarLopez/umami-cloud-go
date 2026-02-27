package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type EnvConfig struct {
	VpcName string  `json:"vpcName"`
	Cidr    string  `json:"cidr"`
	MaxAzs  float64 `json:"maxAzs"`
}

func loadConfig() (*EnvConfig, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	filePath := fmt.Sprintf("config/%s.json", env)

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error leyendo el archivo %s: %w", filePath, err)
	}
	var config EnvConfig
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return nil, fmt.Errorf("Error parseando el Json: %w", err)
	}
	return &config, nil
}

// main entry point
func main() {
	defer jsii.Close()

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)

	envName := os.Getenv("ENV")
	if envName == "" {
		envName = "dev"
	}
	stackName := fmt.Sprintf("%sNetworkStack", envName)

	NewNetworkStack(app, stackName, &NetworkStackProps{
		StackProps: awscdk.StackProps{Env: env()},
		MaxAZs:     config.MaxAzs,
		VpcName:    config.VpcName,
		Cidr:       config.Cidr,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
