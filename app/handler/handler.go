package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		StatusCode: 200,
	}, nil
}
