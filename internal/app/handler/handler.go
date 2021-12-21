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
	d   dao.Dao
	err error
)

func init() {
	d, err = dao.New()
	if err != nil {
		log.Fatalln("ERROR: dynamo connection error", err)
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "http://localhost:8080",
		"Access-Control-Allow-Methods": "GET,OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	}

	var result string
	repository := d.Article()

	switch request.Path {
	case "/article/list":
		switch request.HTTPMethod {
		case "GET":
			var articles []*object.Article
			if articles, err = repository.ListArticles(ctx); err != nil {
				break
			}

			if len(articles) == 0 {
				break
			}

			// 登録日基準で降順に安定ソート
			sort.SliceStable(articles, func(i, j int) bool {
				return pkg.StringToTime(articles[i].CreateTimeStamp).After(pkg.StringToTime(articles[j].CreateTimeStamp))
			})

			result, err = pkg.InterfaceToJson(articles)
		}
	default:
		switch request.HTTPMethod {
		case "GET":
			articleID := request.PathParameters["articleID"]
			if articleID == "" || articleID == "0" {
				return events.APIGatewayProxyResponse{
					Body:       "Not Article Data",
					Headers:    headers,
					StatusCode: http.StatusNotFound,
				}, nil
			}

			var article *object.Article
			if article, err = repository.GetArticle(ctx, articleID); err != nil {
				break
			} else if article.ID == 0 {
				return events.APIGatewayProxyResponse{
					Body:       "Not Article Data",
					Headers:    headers,
					StatusCode: http.StatusNotFound,
				}, nil
			}

			result, err = pkg.InterfaceToJson(article)
		}
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       err.Error(),
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       result,
		Headers:    headers,
		StatusCode: http.StatusOK,
	}, nil
}
