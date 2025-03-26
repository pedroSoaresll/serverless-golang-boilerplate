package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function for Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: fmt.Sprintf(
			"Hello, %s! This is the path: %s",
			request.QueryStringParameters["name"],
			request.PathParameters["name"]),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
