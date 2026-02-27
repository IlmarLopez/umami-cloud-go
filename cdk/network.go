package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NetworkStackProps defines the custom properties for our VPC environments.
type NetworkStackProps struct {
	awscdk.StackProps
	MaxAZs  float64
	VpcName string
	Cidr    string
}

// NewNetworkStack creates the network infrastructure based on provided properties.
func NewNetworkStack(scope constructs.Construct, id string, props *NetworkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Effective Go practice: Indentation style for long struct initializations
	awsec2.NewVpc(stack, jsii.String("VPC"), &awsec2.VpcProps{
		VpcName:     jsii.String(props.VpcName),
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String(props.Cidr)),
		MaxAzs:      jsii.Number(props.MaxAZs),
		NatGateways: jsii.Number(0), // Recommended to avoid unexpected AWS costs
	})

	return stack
}
