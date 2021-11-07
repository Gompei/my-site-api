package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Gompei/my-site-api/app/dao"
	"github.com/aws/aws-lambda-go/events"
)

type App struct {
	dao dao.Dao
}

var (
	err error
	app *App
)

func init() {
	app.dao, err = dao.New()
	if err != nil {
		log.Fatalln("ERROR: dynamo connection error ", err)
	}
}

func Router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                    "application/json",
		"Access-Control-Allow-Origin":     request.Headers["origin"],
		"Access-Control-Allow-Methods":    "GET,POST,PUT,DELETE,OPTIONS",
		"Access-Control-Allow-Headers":    "Origin,Authorization,Accept,X-Requested-With",
		"Access-Control-Allow-Credential": "true",
	}

	var result string

	switch request.HTTPMethod {
	case "GET":
		getHandler(ctx, request)
	case "POST":
	case "PUT":
		putHandler()
	case "DELETE":
	default:
		return events.APIGatewayProxyResponse{
			Headers:    headers,
			Body:       "Not Implemented",
			StatusCode: http.StatusNotImplemented,
		}, nil
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Internal Server Error",
		}, err
	}

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       result,
		StatusCode: http.StatusOK,
	}, nil
}

func getHandler(ctx context.Context, request events.APIGatewayProxyRequest) {

}

func postHandler() {

}

func putHandler() {}

func deleteHandler() {}
