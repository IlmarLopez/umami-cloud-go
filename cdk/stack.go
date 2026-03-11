package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func NewBaseStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = *props
	}
	return awscdk.NewStack(scope, &id, &sprops)
}
