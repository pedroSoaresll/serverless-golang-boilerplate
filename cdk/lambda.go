package main

import (
	"path/filepath"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"github.com/pedroSoaresll/serverless-golang-boilerplate/constants"
)

type APIEndpointConfig struct {
	FunctionName string
	Timeout      int
	MemorySize   int
	Path         string
	Method       awsapigatewayv2.HttpMethod
	Source       string
}

var endpointConfigs = []APIEndpointConfig{
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
		buildEndpoint(stack, httpApi, config)
	}
}

func buildEndpoint(stack awscdk.Stack, httpApi awsapigatewayv2.HttpApi, config APIEndpointConfig) {
	lambdaFunctionId := constants.ENV + config.FunctionName + "LambdaHandler"
	functionProps := &awslambda.FunctionProps{
		Runtime:    awslambda.Runtime_PROVIDED_AL2023(),
		Handler:    jsii.String(filepath.Join(config.Source, "/bootstrap")),
		Code:       awslambda.Code_FromAsset(jsii.String(filepath.Join(config.Source, "function.zip")), nil),
		MemorySize: jsii.Number(float64(config.MemorySize)),
		Timeout:    awscdk.Duration_Seconds(jsii.Number(float64(config.Timeout))),
	}
	lambdaFunction := awslambda.NewFunction(stack, jsii.String(lambdaFunctionId), functionProps)

	integrationId := constants.ENV + config.FunctionName + "LambdaIntegration"
	integrationProps := &awsapigatewayv2integrations.HttpLambdaIntegrationProps{
		Timeout: awscdk.Duration_Seconds(jsii.Number(float64(config.Timeout))),
	}
	integration := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		jsii.String(integrationId),
		lambdaFunction,
		integrationProps,
	)

	addRoutesOptions := &awsapigatewayv2.AddRoutesOptions{
		Path: jsii.String(config.Path),
		Methods: &[]awsapigatewayv2.HttpMethod{
			config.Method,
		},
		Integration: integration,
	}
	httpApi.AddRoutes(addRoutesOptions)
}
