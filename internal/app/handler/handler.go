package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

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
		log.Fatalln("ERROR: dynamo connection error", err)
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "http://localhost:8080",
		"Access-Control-Allow-Methods": "GET,PUT,DELETE,OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	}

	var result string
	repository := d.Article()

	switch request.Path {
	case "/test":
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    headers,
			Body:       fmt.Sprintln("Hello World!!"),
		}, nil

	case "/article/search":
		log.Println("/article/search unimplemented")

	case "/article/list":
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

			result, err = pkg.InterfaceToJson(articles)
		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Implemented",
				StatusCode: http.StatusNotImplemented,
			}, nil
		}

	case "/article":
		switch request.HTTPMethod {
		case "PUT":
			var article *object.Article
			if article, err = object.ToArticleStruct(request.Body); err != nil {
				break
			}

			// TODO
			article.ID, err = repository.GetCountID(ctx)
			if err != nil {
				break
			}
			article.ID++
			err = repository.PutArticle(ctx, article)
			if err != nil {
				break
			}
			err = repository.PutCountID(ctx, article.ID)
			if err != nil {
				break
			}

			result = "Success PUT Article Data"
			articles = nil
		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Implemented",
				StatusCode: http.StatusNotImplemented,
			}, nil
		}

	default:
		switch request.HTTPMethod {
		case "GET":
			if _, err = strconv.Atoi(request.PathParameters["articleID"]); err != nil {
				break
			}

			var article *object.Article
			if article, err = repository.GetArticle(ctx, request.PathParameters["articleID"]); err != nil {
				break
			} else if article.ID == 0 {
				return events.APIGatewayProxyResponse{
					Body:       "Not Found",
					Headers:    headers,
					StatusCode: http.StatusNotFound,
				}, nil
			}
			result, err = pkg.InterfaceToJson(article)

		case "DELETE":
			if _, err = strconv.Atoi(request.PathParameters["articleID"]); err != nil {
				break
			}

			if request.QueryStringParameters["physical"] == "true" && request.PathParameters["articleID"] != "" {
				if err = repository.DeleteArticle(ctx, request.PathParameters["articleID"]); err != nil {
					break
				}
				result = "Success Physical Delete Article Data"
			} else {
				var article *object.Article
				if article, err = object.ToArticleStruct(request.Body); err != nil {
					break
				}
				article.DeleteFlg = true
				err = repository.PutArticle(ctx, article)
				result = "Success Delete Article Data"
			}
			articles = nil

		default:
			return events.APIGatewayProxyResponse{
				Headers:    headers,
				Body:       "Not Found",
				StatusCode: http.StatusNotFound,
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
		Headers:    headers,
		StatusCode: http.StatusOK,
	}, nil
}
