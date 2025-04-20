package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/aws/jsii-runtime-go"
	"github.com/pedroSoaresll/serverless-golang-boilerplate/constants"
)

type ServerlessApiStackProps struct {
	awscdk.StackProps
}

func main() {
	app := awscdk.NewApp(nil)
	stage := constants.ENV

	NewServerlessCognitoStack(app, stage+"ServerlessCognitoStack", &CognitoStackProps{
		awscdk.StackProps{
			StackName: jsii.String(stage + "ServerlessCognitoStack"),
			Env:       env(),
		},
	})

	NewServerlessApiStack(app, stage+"ServerlessApiStack", &ServerlessApiStackProps{
		awscdk.StackProps{
			StackName: jsii.String(stage + "ServerlessGolangStack"),
			Env:       env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(constants.CDK_DEFAULT_ACCOUNT),
		Region:  jsii.String(constants.CDK_DEFAULT_REGION),
	}
}
