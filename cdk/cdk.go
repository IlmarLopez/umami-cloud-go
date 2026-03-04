package main

import (
	"fmt"
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

	envName := os.Getenv("ENV")
	if envName == "" {
		envName = "dev"
	}

	stackName := fmt.Sprintf("umami-stack-%s", envName)

	NewNetworkStack(app, stackName, &NetworkStackProps{
		StackProps: awscdk.StackProps{Env: env()},
		EnvName:    envName,
		Config:     cfg,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil

}
