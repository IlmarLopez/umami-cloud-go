package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func resolveInstanceType(size string) (awsec2.InstanceType, error) {
	if size == "" {
		return nil, fmt.Errorf("The instance size cannot be empty")
	}
	return awsec2.NewInstanceType(jsii.String(size)), nil
}

func NewEC2(stack awscdk.Stack, vpc awsec2.IVpc, props *UmamiStackProps) awsec2.Instance {
	var instanceType awsec2.InstanceType
	instType, err := resolveInstanceType(props.Config.EC2.InstanceSize)
	if err != nil {
		panic(fmt.Sprintf("EC2 configuration error: %v", err))
	} else {
		instanceType = instType
	}

	machineImage := awsec2.MachineImage_LatestAmazonLinux2023(&awsec2.AmazonLinux2023ImageSsmParameterProps{
		CpuType: awsec2.AmazonLinuxCpuType_X86_64,
	})

	instanceName := fmt.Sprintf("%s-Server-%s", props.StackName, props.EnvValue)

	scriptBytes, err := os.ReadFile("scripts/user-data.sh")
	if err != nil {
		panic(fmt.Sprintf("Error crítico: No se encontró scripts/user-data.sh: %v", err))
	}

	userData := awsec2.UserData_Custom(jsii.String(string(scriptBytes)))

	server := awsec2.NewInstance(stack, jsii.String("UmamiServer"), &awsec2.InstanceProps{
		Vpc:          vpc,
		InstanceType: instanceType,
		MachineImage: machineImage,
		InstanceName: jsii.String(instanceName),
		UserData:     userData,
	})

	awscdk.Tags_Of(server).Add(jsii.String("Environment"), jsii.String(props.EnvValue), &awscdk.TagProps{})

	server.Connections().AllowFromAnyIpv4(awsec2.Port_Tcp(jsii.Number(3000)), jsii.String("Allow Umami Web Traffic"))

	server.Connections().AllowFromAnyIpv4(awsec2.Port_Tcp(jsii.Number(22)), jsii.String("Allow SSH Access"))

	return server
}
