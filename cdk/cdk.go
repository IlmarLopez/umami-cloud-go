package main

import (
	"os"

	"umami-cloud-go/cdk/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)

	envValue := os.Getenv("ENV")
	if envValue == "" {
		envValue = "dev"
	}

	stackName := "umami-cloud-go"

	NewVpcStack(app, stackName, &VpcStackProps{
		StackProps: awscdk.StackProps{Env: env()},
		EnvValue:   envValue,
		StackName:  stackName,
		Config:     cfg,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
