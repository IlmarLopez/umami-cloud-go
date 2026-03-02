package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NetworkStackProps defines the custom properties for our VPC environments.
type NetworkStackProps struct {
	awscdk.StackProps
	EnvName string
	Config  *EnvConfig
}

func NewNetworkStack(scope constructs.Construct, id string, props *NetworkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	cidr := "10.0.0.0/16"
	if props.EnvName == "prod" {
		cidr = "10.1.0.0/16"

	}

	vpcName := fmt.Sprintf("UmamiStack-Network-%s", props.EnvName)
	awsec2.NewVpc(stack, jsii.String("VPC"), &awsec2.VpcProps{
		VpcName:     jsii.String(vpcName),
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String(cidr)),
		MaxAzs:      jsii.Number(props.Config.MaxAzs),
		NatGateways: jsii.Number(0),
	})

	return stack
}
