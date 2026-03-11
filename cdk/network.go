package main

import (
	"fmt"
	"umami-cloud-go/cdk/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// NetworkStackProps defines the custom properties for our VPC environments.
type VpcStackProps struct {
	awscdk.StackProps
	EnvValue  string
	StackName string
	Config    *config.EnvConfig
}

func NewVpcStack(scope constructs.Construct, id string, props *VpcStackProps) awscdk.Stack {

	stack := NewBaseStack(scope, id, &props.StackProps)

	vpcName := fmt.Sprintf("%s-Network-%s", props.StackName, props.EnvValue)

	awsec2.NewVpc(stack, jsii.String("VPC"), &awsec2.VpcProps{
		VpcName:     jsii.String(vpcName),
		MaxAzs:      jsii.Number(props.Config.VPC.MaxAZs),
		NatGateways: jsii.Number(0),
	})

	return stack
}
