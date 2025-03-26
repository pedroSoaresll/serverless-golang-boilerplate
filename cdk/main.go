package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/pedroSoaresll/serverless-golang-boilerplate/constants"
)

type ServerlessApiStackProps struct {
	awscdk.StackProps
}

func NewServerlessApiStack(scope constructs.Construct, id string, props *ServerlessApiStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	httpApi := awsapigatewayv2.NewHttpApi(stack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiProps{
		ApiName: jsii.String(constants.ENV + "HttpApi"),
	})

	LoadApiGatewayEndpoints(stack, httpApi)

	// Output the HTTP API URL
	awscdk.NewCfnOutput(stack, jsii.String(constants.ENV+"HttpApiUrl"), &awscdk.CfnOutputProps{
		Value:       httpApi.Url(),
		Description: jsii.String("The base URL for the HTTP API Gateway"),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)
	stage := constants.ENV

	NewServerlessApiStack(app, stage+"ServerlessApiStack", &ServerlessApiStackProps{
		awscdk.StackProps{
			StackName: jsii.String(stage + "ServerlessGolangStack"),
			Env:       env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(constants.CDK_DEFAULT_ACCOUNT),
		Region:  jsii.String(constants.CDK_DEFAULT_REGION),
	}
}
