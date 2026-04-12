package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func NewVPC(stack awscdk.Stack, id string, props *UmamiStackProps) awsec2.Vpc {

	vpcName := fmt.Sprintf("%s-Network-%s", props.StackName, props.EnvValue)

	return awsec2.NewVpc(stack, jsii.String("VPC"), &awsec2.VpcProps{
		VpcName:     jsii.String(vpcName),
		MaxAzs:      jsii.Number(2),
		NatGateways: jsii.Number(0),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				Name:       jsii.String("PublicSubnet"),
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
		},
	})

}
