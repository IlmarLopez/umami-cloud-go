package main

import (
	"umami-cloud-go/pkg/cfg"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UmamiCloudGoStackProps struct {
	awscdk.StackProps
	Config cfg.EnvConfig
}

func NewUmamiCloudGoStack(scope constructs.Construct, id string, props *UmamiCloudGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// La VPC que tenías en tu imagen, ahora integrada aquí
	awsec2.NewVpc(stack, jsii.String("UmamiVPC"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
		MaxAzs:      jsii.Number(2),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				Name:       jsii.String("Public"),
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
		},
	})

	return stack
}

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)

	// Stack de Desarrollo
	devCfg := cfg.GetDevConfig()
	NewUmamiCloudGoStack(app, "UmamiStackDev", &UmamiCloudGoStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{Account: jsii.String(devCfg.Account), Region: jsii.String(devCfg.Region)},
		},
		Config: devCfg,
	})

	// Stack de Producción
	prodCfg := cfg.GetProdConfig()
	NewUmamiCloudGoStack(app, "UmamiStackProd", &UmamiCloudGoStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{Account: jsii.String(prodCfg.Account), Region: jsii.String(prodCfg.Region)},
		},
		Config: prodCfg,
	})

	app.Synth(nil)
}
