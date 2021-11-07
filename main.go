package main

import (
	"github.com/Gompei/my-site-api/app/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
