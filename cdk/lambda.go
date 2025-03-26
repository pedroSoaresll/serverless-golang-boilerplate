package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/pedroSoaresll/serverless-golang-boilerplate/helpers"
)

var endpointConfigs = []helpers.APIEndpointConfig{
	{
		Source:       "lambda",
		FunctionName: "Hello",
		Path:         "/hello",
		Method:       awsapigatewayv2.HttpMethod_GET,
		Timeout:      5,
		MemorySize:   128,
	},
	{
		Source:       "lambda/world",
		FunctionName: "HelloWorld",
		Path:         "/hello/{name}",
		Method:       awsapigatewayv2.HttpMethod_GET,
		Timeout:      5,
		MemorySize:   128,
	},
}

func LoadApiGatewayEndpoints(stack awscdk.Stack, httpApi awsapigatewayv2.HttpApi) {
	for _, config := range endpointConfigs {
		helpers.NewAPIEndpoint(stack, httpApi, config)
	}
}
