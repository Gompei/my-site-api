package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/Gompei/my-site-api/app/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda") {
		lambda.Start(handler.Handler)
	} else {
		response, err := run()
		if err != nil {
			log.Fatalf("ERROR handler %v\n", err)
		}
		log.Printf("INFO handler %v\n", response)
	}
}

func run() (map[string]interface{}, error) {
	ctx := context.Background()
	event := events.APIGatewayProxyRequest{
		Resource: "/api",
		Path:     "/api",
		Headers: map[string]string{
			"Host":      "example.com",
			"x-api-key": "abc",
		},
		MultiValueHeaders: nil,
		QueryStringParameters: map[string]string{
			"test": "1",
		},
	}

	httpMethods := []string{"GET", "POST", "PUT", "DELETE"}
	result := make(map[string]interface{}, len(httpMethods))
	for _, httpMethod := range httpMethods {
		event.HTTPMethod = httpMethod
		response, err := handler.Handler(ctx, event)
		if err != nil {
			return nil, err
		}
		result[httpMethod] = response
	}

	return result, nil
}
