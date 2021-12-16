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
		callAPIs := []func() (map[string]interface{}, error){
			runArticleAPI,
			runArticleListAPI,
		}

		for _, api := range callAPIs {
			response, err := api()
			if err != nil {
				log.Fatalf("ERROR %v\n", err)
			}
			log.Printf("INFO %v\n", response)
		}
	}
}

// Article CRUD API
func runArticleAPI() (map[string]interface{}, error) {
	ctx := context.Background()
	event := events.APIGatewayProxyRequest{
		Path: "/article",
		Headers: map[string]string{
			"Host":      "example.com",
			"x-api-key": "abc",
		},
		PathParameters: map[string]string{
			"articleID": "101",
		},
	}

	article := object.Article{
		ID:              101,
		Title:           "example",
		SubTitle:        "example",
		ImageURL:        "example",
		CategoryTag:     []string{"example"},
		Content:         "example",
		CreateTimeStamp: pkg.CreateTimeStamp(),
		UpdateTimeStamp: pkg.CreateTimeStamp(),
		PublicFlg:       false,
		DeleteFlg:       false,
	}
	j, err := pkg.InterfaceToJson(article)
	if err != nil {
		return nil, err
	}
	event.Body = j

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

// Article List API
func runArticleListAPI() (map[string]interface{}, error) {
	ctx := context.Background()
	event := events.APIGatewayProxyRequest{
		Resource: "/article/list",
		Path:     "/article/list",
		Headers: map[string]string{
			"Host":      "example.com",
			"x-api-key": "abc",
		},
	}

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
