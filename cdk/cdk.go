package main

import (
	"fmt"
	"os"

	"umami-cloud-go/cdk/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Note: .env file not found, using system environment variables.")
	}

	app := awscdk.NewApp(&awscdk.AppProps{})

	defer jsii.Close()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	envValue := os.Getenv("ENV")
	if envValue == "" {
		envValue = "dev"
	}

	stackName := "umami-cloud-go"

	NewUmamiCloudGoStack(app, stackName, &UmamiStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
		EnvValue:  envValue,
		StackName: stackName,
		Config:    cfg,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	account := os.Getenv("CDK_DEFAULT_ACCOUNT")
	region := os.Getenv("CDK_DEFAULT_REGION")

	if account == "" || region == "" {
		return nil
	}

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
