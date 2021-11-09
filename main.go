package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/Gompei/my-site-api/app/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda") {
		lambda.Start(handler.Handler)
	} else {
		ctx := context.Background()
		aa := events.APIGatewayProxyRequest{
			Resource: "aa",
		}
		response, err := handler.Handler(ctx, aa)
		if err != nil {
			log.Fatalf("ERROR handler %v\n", err)
		}
		log.Printf("INFO handler %v\n", response)
	}
}
