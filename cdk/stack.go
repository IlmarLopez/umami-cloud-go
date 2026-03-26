package main

import (
	"umami-cloud-go/cdk/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type UmamiStackProps struct {
	awscdk.StackProps
	EnvValue  string
	StackName string
	Config    *config.EnvConfig
}

func NewUmamiCloudGoStack(scope constructs.Construct, id string, props *UmamiStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := NewVPC(stack, id, props)

	NewEC2(stack, vpc, props)
	return stack
}
