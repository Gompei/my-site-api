package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Gompei/my-site-api/app/dao"
	"github.com/aws/aws-lambda-go/events"
)

type App struct {
	dao dao.Dao
}

type GetArticleRequest struct {
	articleID int64 `json:"article_id,omitempty"`
	limit     int64 `json:"limit,omitempty"`
}

var (
	err error
	app *App
)

//func init() {
//	app.dao, err = dao.New()
//	if err != nil {
//		log.Fatalln("ERROR: dynamo connection error ", err)
//	}
//}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//headers := map[string]string{
	//	"Content-Type":                    "application/json",
	//	"Access-Control-Allow-Origin":     request.Headers["origin"],
	//	"Access-Control-Allow-Methods":    "GET,POST,PUT,DELETE,OPTIONS",
	//	"Access-Control-Allow-Headers":    "Origin,Authorization,Accept,X-Requested-With",
	//	"Access-Control-Allow-Credential": "true",
	//}
	//
	//var result string
	//var article *object.Article
	//repository := app.dao.Article()
	//
	//switch request.HTTPMethod {
	//case "GET":
	//	var req GetArticleRequest
	//	err = json.Unmarshal([]byte(request.Body), &req)
	//	if err != nil {
	//		break
	//	}
	//	if req.articleID == 0 {
	//		var articles []*object.Article
	//		if articles, err = repository.GetAllArticle(); err != nil {
	//			break
	//		}
	//		result, err = toJson(articles)
	//	} else {
	//		if article, err = repository.GetArticle(req.articleID); err != nil {
	//			break
	//		}
	//		result, err = toJson(article)
	//	}
	//case "POST", "PUT":
	//	if article, err = object.ToArticleStruct(request.Body); err != nil {
	//		break
	//	}
	//	err = repository.PutArticle(article)
	//	result = "Success POST・PUT Article Data"
	//case "DELETE":
	//	article, err = object.ToArticleStruct(request.Body)
	//	if err != nil {
	//		break
	//	}
	//	err = repository.DeleteArticle(article)
	//	result = "Success DELETE Article Data"
	//default:
	//	return events.APIGatewayProxyResponse{
	//		Headers:    headers,
	//		Body:       "Not Implemented",
	//		StatusCode: http.StatusNotImplemented,
	//	}, nil
	//}
	//
	//if err != nil {
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: http.StatusInternalServerError,
	//		Body:       err.Error(),
	//	}, err
	//}

	return events.APIGatewayProxyResponse{
		Body:       "test",
		StatusCode: http.StatusOK,
	}, nil
}

func toJson(s interface{}) (string, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(j), err
}
