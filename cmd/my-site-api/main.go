package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/Gompei/my-site-api/internal/app/domain/object"
	"github.com/Gompei/my-site-api/internal/app/handler"
	"github.com/Gompei/my-site-api/pkg"
	"github.com/aws/aws-lambda-go/events"
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
		Path:     "/api/article/list",
		Headers: map[string]string{
			"Host":      "example.com",
			"x-api-key": "abc",
		},
		PathParameters: map[string]string{
			"articleID": "111",
		},
	}

	article := object.Article{
		ID:              1,
		Title:           "example",
		SubTitle:        "example",
		ImageURL:        "example",
		CategoryTag:     "example",
		Description:     "example",
		Content:         "example",
		CreateTimeStamp: pkg.CreateTimeStamp(),
		UpdateTimeStamp: pkg.CreateTimeStamp(),
		DeleteFlg:       false,
	}
	j, err := pkg.InterfaceToJson(article)
	if err != nil {
		return nil, err
	}
	event.Body = j

	//httpMethods := []string{"GET", "POST", "PUT", "DELETE"}
	httpMethods := []string{"GET"}
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
