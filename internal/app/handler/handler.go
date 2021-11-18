package handler

import (
	"context"
	"log"
	"net/http"
	"sort"

	"github.com/Gompei/my-site-api/internal/app/dao"
	"github.com/Gompei/my-site-api/internal/app/domain/object"
	"github.com/Gompei/my-site-api/pkg"
	"github.com/aws/aws-lambda-go/events"
)

var (
	d        dao.Dao
	err      error
	articles []*object.Article
)

func init() {
	d, err = dao.New()
	if err != nil {
		log.Fatalln("ERROR: dynamo connection error ", err)
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                    "application/json",
		"Access-Control-Allow-Origin":     request.Headers["origin"],
		"Access-Control-Allow-Methods":    "GET,POST,PUT,DELETE,OPTIONS",
		"Access-Control-Allow-Headers":    "Origin,Authorization,Accept,X-Requested-With",
		"Access-Control-Allow-Credential": "true",
	}

	var result string
	repository := d.Article()

	switch request.Path {
	case "/api/article/search":
		log.Println("/api/article/search unimplemented")
	case "/api/article/list":
		switch request.HTTPMethod {
		case "GET":
			if articles == nil {
				if articles, err = repository.ListArticles(ctx); err != nil {
					break
				}

				// 登録日基準で降順に安定ソート
				sort.SliceStable(articles, func(i, j int) bool {
					return pkg.StringToTime(articles[i].CreateTimeStamp).After(pkg.StringToTime(articles[j].CreateTimeStamp))
				})
			}

			// TODO:入力値チェックをAPI Gateway側で実施
			page := pkg.StringToInt(request.QueryStringParameters["p"])
			result, err = pkg.InterfaceToJson(articles[:page])
		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Implemented",
				StatusCode: http.StatusNotImplemented,
			}, nil
		}
	case "/api/article/physical-delete":
		switch request.HTTPMethod {
		case "DELETE":
			id := request.PathParameters["articleID"]
			err = repository.DeleteArticle(ctx, id)
			result = "Success Physical DELETE Article Data"
			articles = nil
		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Implemented",
				StatusCode: http.StatusNotImplemented,
			}, nil
		}
	case "/api/article":
		switch request.HTTPMethod {
		case "GET":
			id := request.PathParameters["articleID"]
			var article *object.Article
			if article, err = repository.GetArticle(ctx, id); err != nil {
				break
			}
			result, err = pkg.InterfaceToJson(article)
		case "POST", "PUT":
			var article *object.Article
			if article, err = object.ToArticleStruct(request.Body); err != nil {
				break
			}
			err = repository.PutArticle(ctx, article)
			result = "Success POST・PUT Article Data"
			articles = nil
		case "DELETE":
			var article *object.Article
			if article, err = object.ToArticleStruct(request.Body); err != nil {
				break
			}
			article.DeleteFlg = true
			err = repository.PutArticle(ctx, article)
			result = "Success Logic DELETE Article Data"
			articles = nil
		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Implemented",
				StatusCode: http.StatusNotImplemented,
			}, nil
		}
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       result,
		StatusCode: http.StatusOK,
	}, nil
}
