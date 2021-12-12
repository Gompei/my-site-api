package dao

import (
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamo() (*dynamodb.DynamoDB, error) {
	sslFlg := os.Getenv("DISABLE_SSL_FLG")
	if sslFlg == "" {
		sslFlg = "false"
	}

	disableSSLFlg, err := strconv.ParseBool(sslFlg)
	if err != nil {
		return nil, err
	}

	region := os.Getenv("REGION")
	if region == "" {
		region = "ap-northeast-1"
	}

	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(region),
		Endpoint:   aws.String(os.Getenv("ENDPOINT")),
		DisableSSL: aws.Bool(disableSSLFlg),
	}))), nil
}
