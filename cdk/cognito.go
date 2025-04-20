package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/pedroSoaresll/serverless-golang-boilerplate/constants"
)

type CognitoStackProps struct {
	awscdk.StackProps
}

func NewServerlessCognitoStack(scope constructs.Construct, id string, props *CognitoStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	userPool := awscognito.NewUserPool(stack, jsii.String(constants.ENV+"UserPool"), &awscognito.UserPoolProps{
		UserPoolName:        jsii.String(constants.ENV + "TestUserPool"),
		SignInCaseSensitive: jsii.Bool(false),
		SelfSignUpEnabled:   jsii.Bool(true),
		UserVerification: &awscognito.UserVerificationConfig{
			EmailSubject: jsii.String("Seu código de confirmação"),
			EmailBody:    jsii.String("Seu código de confirmação é {####}"),
			EmailStyle:   awscognito.VerificationEmailStyle_CODE,
			SmsMessage:   jsii.String("Seu código de confirmação é {####}"),
		},
		UserInvitation: &awscognito.UserInvitationConfig{
			EmailSubject: jsii.String("Convite para se juntar ao nosso aplicativo"),
			EmailBody:    jsii.String("Oi {username}, você foi convidado a se juntar ao nosso aplicativo! Sua senha temporária é {####}."),
			SmsMessage:   jsii.String("Oi {username}, você foi convidado a se juntar ao nosso aplicativo! Sua senha temporária é {####}."),
		},
		SignInAliases: &awscognito.SignInAliases{
			Email:    jsii.Bool(true),
			Phone:    jsii.Bool(true),
			Username: jsii.Bool(true),
		},
		AutoVerify: &awscognito.AutoVerifiedAttrs{
			Email: jsii.Bool(true),
			Phone: jsii.Bool(true),
		},
		StandardAttributes: &awscognito.StandardAttributes{
			GivenName: &awscognito.StandardAttribute{
				Required: jsii.Bool(true),
				Mutable:  jsii.Bool(true),
			},
			FamilyName: &awscognito.StandardAttribute{
				Required: jsii.Bool(true),
				Mutable:  jsii.Bool(true),
			},
		},
		CustomAttributes: &map[string]awscognito.ICustomAttribute{
			"role": awscognito.NewStringAttribute(&awscognito.StringAttributeProps{
				Mutable: jsii.Bool(true),
			}),
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String(constants.ENV+"UserPoolId"), &awscdk.CfnOutputProps{
		Value:       userPool.UserPoolId(),
		Description: jsii.String("The ID of the User Pool"),
	})

	return stack
}
